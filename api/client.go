package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strconv"

	"github.com/ssoroka/slice"

	"github.com/infrahq/infra/internal/logging"
	"github.com/infrahq/infra/uid"
)

var apiVersion = "0.13.0"

var ErrTimeout = errors.New("client timed out waiting for response from server")

const (
	InfraAdminRole     = "admin"
	InfraViewRole      = "view"
	InfraConnectorRole = "connector"
)

type Client struct {
	Name      string
	Version   string
	URL       string
	AccessKey string
	HTTP      http.Client
	// Headers are HTTP headers that will be added to every request made by the Client.
	Headers http.Header
}

// checkError checks the resp for an error code, and returns an api.Error with
// details about the error. Returns nil if the status code is 2xx.
//
// 3xx codes are considered an error because redirects should have already
// been followed before calling checkError.
func checkError(req *http.Request, resp *http.Response, body []byte) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	apiError := Error{
		Method: req.Method,
		Path:   req.URL.Path,
		Code:   int32(resp.StatusCode),
	}

	err := json.Unmarshal(body, &apiError)
	if err != nil {
		// Use the full body as the message if we fail to decode a response.
		apiError.Message = string(body)
	}

	return apiError
}

// ErrorStatusCode returns the http status code from the error.
// Returns 0 if the error is nil, or if the error is not of type Error.
func ErrorStatusCode(err error) int32 {
	var apiError Error
	if errors.As(err, &apiError) {
		return apiError.Code
	}
	return 0
}

func request[Req, Res any](client Client, method string, path string, query Query, reqBody *Req) (*Res, error) {
	var body []byte

	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("marshal json: %w", err)
		}

		body = b
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", client.URL, path), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = url.Values(query).Encode()

	clientName, clientVersion := "client", "unknown"
	if client.Name != "" {
		clientName = client.Name
	}

	if client.Version != "" {
		clientVersion = client.Version
	}

	req.Header.Add("Authorization", "Bearer "+client.AccessKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Infra-Version", apiVersion)
	req.Header.Set("User-Agent", fmt.Sprintf("Infra/%v (%s %v; %v/%v)", apiVersion, clientName, clientVersion, runtime.GOOS, runtime.GOARCH))

	for k, v := range client.Headers {
		req.Header[k] = v
	}

	resp, err := client.HTTP.Do(req)
	if err != nil {
		if connError := HandleConnError(err); connError != nil {
			return nil, connError
		}
		return nil, fmt.Errorf("%s %q: %w", method, path, err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("%w: %s", ErrTimeout, err)
		}
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if err := checkError(req, resp, body); err != nil {
		return nil, err
	}

	var resBody Res
	if len(body) > 0 {
		if err := json.Unmarshal(body, &resBody); err != nil {
			return nil, fmt.Errorf("parsing json response: %w. partial text: %q", err, partialText(body, 100))
		}
	}

	return &resBody, nil
}

func get[Res any](client Client, path string, query Query) (*Res, error) {
	return request[EmptyRequest, Res](client, http.MethodGet, path, query, nil)
}

func post[Req, Res any](client Client, path string, req *Req) (*Res, error) {
	return request[Req, Res](client, http.MethodPost, path, Query{}, req)
}

func put[Req, Res any](client Client, path string, req *Req) (*Res, error) {
	return request[Req, Res](client, http.MethodPut, path, Query{}, req)
}

func patch[Req, Res any](client Client, path string, req *Req) (*Res, error) {
	return request[Req, Res](client, http.MethodPatch, path, Query{}, req)
}

func delete(client Client, path string) error {
	_, err := request[EmptyRequest, EmptyResponse](client, http.MethodDelete, path, Query{}, nil)
	return err
}

func (c Client) ListUsers(req ListUsersRequest) (*ListResponse[User], error) {
	ids := slice.Map[uid.ID, string](req.IDs, func(id uid.ID) string {
		return id.String()
	})
	return get[ListResponse[User]](c, "/api/users", Query{"name": {req.Name}, "group": {req.Group.String()}, "ids": ids})
}

func (c Client) GetUser(id uid.ID) (*User, error) {
	return get[User](c, fmt.Sprintf("/api/users/%s", id), Query{})
}

func (c Client) CreateUser(req *CreateUserRequest) (*CreateUserResponse, error) {
	return post[CreateUserRequest, CreateUserResponse](c, "/api/users", req)
}

func (c Client) UpdateUser(req *UpdateUserRequest) (*User, error) {
	return put[UpdateUserRequest, User](c, fmt.Sprintf("/api/users/%s", req.ID.String()), req)
}

func (c Client) DeleteUser(id uid.ID) error {
	return delete(c, fmt.Sprintf("/api/users/%s", id))
}

// Deprecated: use ListGrants
func (c Client) ListUserGrants(id uid.ID) (*ListResponse[Grant], error) {
	return get[ListResponse[Grant]](c, fmt.Sprintf("/api/users/%s/grants", id), Query{})
}

func (c Client) ListGroups(req ListGroupsRequest) (*ListResponse[Group], error) {
	return get[ListResponse[Group]](c, "/api/groups", Query{
		"name":   {req.Name},
		"userID": {req.UserID.String()},
	})
}

func (c Client) GetGroup(id uid.ID) (*Group, error) {
	return get[Group](c, fmt.Sprintf("/api/groups/%s", id), Query{})
}

func (c Client) CreateGroup(req *CreateGroupRequest) (*Group, error) {
	return post[CreateGroupRequest, Group](c, "/api/groups", req)
}

func (c Client) DeleteGroup(id uid.ID) error {
	return delete(c, fmt.Sprintf("/api/groups/%s", id))
}

func (c Client) UpdateUsersInGroup(req *UpdateUsersInGroupRequest) error {
	_, err := patch[UpdateUsersInGroupRequest, EmptyResponse](c, fmt.Sprintf("/api/groups/%s/users", req.GroupID), req)
	return err
}

// Deprecated: use ListGrants
func (c Client) ListGroupGrants(id uid.ID) (*ListResponse[Grant], error) {
	return get[ListResponse[Grant]](c, fmt.Sprintf("/api/groups/%s/grants", id), Query{})
}

func (c Client) ListProviders(name string) (*ListResponse[Provider], error) {
	return get[ListResponse[Provider]](c, "/api/providers", Query{"name": {name}})
}

func (c Client) GetProvider(id uid.ID) (*Provider, error) {
	return get[Provider](c, fmt.Sprintf("/api/providers/%s", id), Query{})
}

func (c Client) CreateProvider(req *CreateProviderRequest) (*Provider, error) {
	return post[CreateProviderRequest, Provider](c, "/api/providers", req)
}

func (c Client) UpdateProvider(req UpdateProviderRequest) (*Provider, error) {
	return put[UpdateProviderRequest, Provider](c, fmt.Sprintf("/api/providers/%s", req.ID.String()), &req)
}

func (c Client) DeleteProvider(id uid.ID) error {
	return delete(c, fmt.Sprintf("/api/providers/%s", id))
}

func (c Client) ListGrants(req ListGrantsRequest) (*ListResponse[Grant], error) {
	return get[ListResponse[Grant]](c, "/api/grants", Query{
		"user":          {req.User.String()},
		"group":         {req.Group.String()},
		"resource":      {req.Resource},
		"privilege":     {req.Privilege},
		"showInherited": {strconv.FormatBool(req.ShowInherited)},
	})
}

func (c Client) CreateGrant(req *CreateGrantRequest) (*CreateGrantResponse, error) {
	return post[CreateGrantRequest, CreateGrantResponse](c, "/api/grants", req)
}

func (c Client) DeleteGrant(id uid.ID) error {
	return delete(c, fmt.Sprintf("/api/grants/%s", id))
}

func (c Client) ListDestinations(req ListDestinationsRequest) (*ListResponse[Destination], error) {
	return get[ListResponse[Destination]](c, "/api/destinations", Query{
		"name":      {req.Name},
		"unique_id": {req.UniqueID},
	})
}

func (c Client) CreateDestination(req *CreateDestinationRequest) (*Destination, error) {
	return post[CreateDestinationRequest, Destination](c, "/api/destinations", req)
}

func (c Client) UpdateDestination(req UpdateDestinationRequest) (*Destination, error) {
	return put[UpdateDestinationRequest, Destination](c, fmt.Sprintf("/api/destinations/%s", req.ID.String()), &req)
}

func (c Client) DeleteDestination(id uid.ID) error {
	return delete(c, fmt.Sprintf("/api/destinations/%s", id))
}

func (c Client) ListAccessKeys(req ListAccessKeysRequest) (*ListResponse[AccessKey], error) {
	return get[ListResponse[AccessKey]](c, "/api/access-keys", Query{
		"user_id":      {req.UserID.String()},
		"name":         {req.Name},
		"show_expired": {fmt.Sprint(req.ShowExpired)},
	})
}

func (c Client) CreateAccessKey(req *CreateAccessKeyRequest) (*CreateAccessKeyResponse, error) {
	return post[CreateAccessKeyRequest, CreateAccessKeyResponse](c, "/api/access-keys", req)
}

func (c Client) DeleteAccessKey(id uid.ID) error {
	return delete(c, fmt.Sprintf("/api/access-keys/%s", id))
}

func (c Client) CreateToken() (*CreateTokenResponse, error) {
	return post[EmptyRequest, CreateTokenResponse](c, "/api/tokens", &EmptyRequest{})
}

func (c Client) Login(req *LoginRequest) (*LoginResponse, error) {
	return post[LoginRequest, LoginResponse](c, "/api/login", req)
}

func (c Client) Logout() error {
	_, err := post[EmptyRequest, EmptyResponse](c, "/api/logout", &EmptyRequest{})
	return err
}

func (c Client) SignupEnabled() (*SignupEnabledResponse, error) {
	return get[SignupEnabledResponse](c, "/api/signup", Query{})
}

func (c Client) Signup(req *SignupRequest) (*CreateAccessKeyResponse, error) {
	return post[SignupRequest, CreateAccessKeyResponse](c, "/api/signup", req)
}

func (c Client) GetServerVersion() (*Version, error) {
	return get[Version](c, "/api/version", Query{})
}

func partialText(body []byte, limit int) string {
	if len(body) <= limit {
		return string(body)
	}

	return string(body[:limit]) + "..."
}

// HandleConnError translates common connection errors into more informative human
// readable errors. Returns `nil` if the error was not handled, so it is the callers responsibility to
// return the original error if `HandleConnError` returns nil.
func HandleConnError(err error) error {
	urlErr := &url.Error{}
	if errors.As(err, &urlErr) {
		if urlErr.Timeout() {
			return fmt.Errorf("%w: %s", ErrTimeout, err)
		}
	}

	if errors.Is(err, io.EOF) {
		logging.Debugf("request error: %v", err)
		return fmt.Errorf("could not reach infra server, please wait a moment and try again")
	}

	return nil
}
