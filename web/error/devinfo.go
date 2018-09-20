package error

// DevInfo contains development information that is passed to error handlers in development mode only.
type DevInfo struct {
	stacktrace []byte
	file       string
	line       int
}
