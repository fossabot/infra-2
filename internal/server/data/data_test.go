package data

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"gorm.io/gorm"
	"gotest.tools/v3/assert"

	"github.com/infrahq/infra/internal/logging"
	"github.com/infrahq/infra/internal/server/models"
	"github.com/infrahq/infra/internal/testing/patch"
	"github.com/infrahq/infra/uid"
)

func setupDB(t *testing.T, driver gorm.Dialector) *gorm.DB {
	t.Helper()
	patch.ModelsSymmetricKey(t)

	db, err := NewDB(driver, nil)
	assert.NilError(t, err)

	InfraProvider(db)

	logging.PatchLogger(t)
	t.Cleanup(InvalidateCache)

	return db
}

// dbDrivers returns the list of database drivers to test.
// Set POSTGRESQL_CONNECTION to a postgresql connection string to run tests
// against postgresql.
func dbDrivers(t *testing.T) []gorm.Dialector {
	t.Helper()
	var drivers []gorm.Dialector

	tmp := t.TempDir()
	sqlite, err := NewSQLiteDriver(filepath.Join(tmp, t.Name()))
	assert.NilError(t, err, "sqlite driver")
	drivers = append(drivers, sqlite)

	pgConn, ok := os.LookupEnv("POSTGRESQL_CONNECTION")
	switch {
	case ok:
		pgsql, err := NewPostgresDriver(pgConn)
		assert.NilError(t, err, "postgresql driver")

		db, err := gorm.Open(pgsql)
		assert.NilError(t, err, "connect to postgresql")
		t.Cleanup(func() {
			assert.NilError(t, db.Exec("DROP SCHEMA testing CASCADE").Error)
		})
		assert.NilError(t, db.Exec("CREATE SCHEMA testing").Error)

		pgsql, err = NewPostgresDriver(pgConn + " search_path=testing")
		assert.NilError(t, err, "postgresql driver")

		drivers = append(drivers, pgsql)
	case os.Getenv("CI") != "":
		t.Fatalf("CI must test all drivers, set POSTGRESQL_CONNECTION")
	}

	return drivers
}

// runDBTests against all supported databases. Defaults to only sqlite locally,
// and all supported DBs in CI. See dbDrivers to test other databases locally.
func runDBTests(t *testing.T, run func(t *testing.T, db *gorm.DB)) {
	for _, driver := range dbDrivers(t) {
		t.Run(driver.Name(), func(t *testing.T) {
			run(t, setupDB(t, driver))
		})
	}
}

func TestSnowflakeIDSerialization(t *testing.T) {
	runDBTests(t, func(t *testing.T, db *gorm.DB) {
		id := uid.New()
		g := &models.Group{Model: models.Model{ID: id}, Name: "Foo"}
		err := db.Create(g).Error
		assert.NilError(t, err)

		var group models.Group
		err = db.First(&group, &models.Group{Name: "Foo"}).Error
		assert.NilError(t, err)
		assert.Assert(t, 0 != group.ID)

		var intID int64
		err = db.Select("id").Table("groups").Scan(&intID).Error
		assert.NilError(t, err)

		assert.Equal(t, int64(id), intID)
	})
}

func TestDatabaseSelectors(t *testing.T) {
	driver, err := NewSQLiteDriver("file::memory:")
	assert.NilError(t, err)

	db, err := newRawDB(driver)
	assert.NilError(t, err)
	t.Logf("DB pointer: %p", db)

	assert.NilError(t, initializeSchema(db))

	// mimic server.DatabaseMiddleware
	withCtx := db.WithContext(context.Background())
	assert.Assert(t, db != withCtx, "db=%p withCtx=%p", db, withCtx)
	t.Logf("DB pointer: %p", withCtx)

	err = withCtx.Transaction(func(tx *gorm.DB) error {
		assert.Assert(t, withCtx != tx, "db=%p tx=%p", withCtx, tx)
		t.Logf("DB pointer: %p", tx)

		// query using one of our helpers and selectors
		_, err := ListGrants(tx, &models.Pagination{}, ByID(534))
		assert.NilError(t, err)

		// query with Model and Where
		var groups []models.Group
		qDB := tx.Model(&models.Group{}).Where("id = ?", 42).Find(&groups)
		assert.NilError(t, qDB.Error)
		assert.Assert(t, tx != qDB, "tx=%p queryDB=%p", tx, qDB)
		t.Logf("DB pointer: %p", qDB)

		// Show that queries have not modified the original gorm.DB references
		assert.Equal(t, len(db.Statement.Clauses), 0)
		assert.Equal(t, len(withCtx.Statement.Clauses), 0)
		assert.Equal(t, len(tx.Statement.Clauses), 0)
		return nil
	})
	assert.NilError(t, err)

	// query using one of our helpers and selectors
	_, err = ListGrants(db, &models.Pagination{}, ByID(534))
	assert.NilError(t, err)

	// query with Model and Where
	var groups []models.Group
	qDB := db.Model(&models.Group{}).Where("id = ?", 42).Find(&groups)
	assert.NilError(t, qDB.Error)
	assert.Assert(t, db != qDB, "db=%p queryDB=%p", db, qDB)
	t.Logf("DB pointer: %p", qDB)

	// Show that queries have not modified the original gorm.DB references
	assert.Equal(t, len(db.Statement.Clauses), 0)
	assert.Equal(t, len(withCtx.Statement.Clauses), 0)
}

func TestPaginationSelector(t *testing.T) {
	letters := make([]string, 0, 26)
	runDBTests(t, func(t *testing.T, db *gorm.DB) {
		for r := 'a'; r < 'a'+26; r++ {
			letters = append(letters, string(r))
			g := &models.Identity{Name: string(r)}
			err := db.Create(g).Error
			assert.NilError(t, err)
		}

		p := models.Pagination{Page: 1, Limit: 10}

		actual, err := ListIdentities(db, &p)
		assert.NilError(t, err)
		assert.Equal(t, len(actual), 10)
		for i := 0; i < p.Limit; i++ {
			assert.Equal(t, letters[i+(p.Page-1)*p.Limit], actual[i].Name)
		}

		p.Page = 2
		actual, err = ListIdentities(db, &p)
		assert.NilError(t, err)
		assert.Equal(t, len(actual), 10)
		for i := 0; i < p.Limit; i++ {
			assert.Equal(t, letters[i+(p.Page-1)*p.Limit], actual[i].Name)
		}

		p.Page = 3
		actual, err = ListIdentities(db, &p)
		assert.NilError(t, err)
		assert.Equal(t, len(actual), 6)

		for i := 0; i < 6; i++ {
			assert.Equal(t, letters[i+(p.Page-1)*p.Limit], actual[i].Name)
		}

		p.Page, p.Limit = 1, 26
		actual, err = ListIdentities(db, &p)
		assert.NilError(t, err)
		for i, user := range actual {
			assert.Equal(t, user.Name, letters[i])
		}

	})
}

func TestDefaultSortFromType(t *testing.T) {
	assert.Equal(t, getDefaultSortFromType(new(models.AccessKey)), "name ASC")
	assert.Equal(t, getDefaultSortFromType(new(models.Destination)), "name ASC")
	assert.Equal(t, getDefaultSortFromType(new(models.Grant)), "id ASC")
	assert.Equal(t, getDefaultSortFromType(new(models.Group)), "name ASC")
	assert.Equal(t, getDefaultSortFromType(new(models.Provider)), "name ASC")
	assert.Equal(t, getDefaultSortFromType(new(models.Identity)), "name ASC")
}
