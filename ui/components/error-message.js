export default function ErrorMessage({ message, center = false }) {
  return (
    <p
      className={`${
        center ? 'mt-2 text-center' : 'mb-1'
      } text-xs text-pink-500`}
    >
      {message}
    </p>
  )
}
