package smarterror

// IError ...
type IError interface {
	Error() string
	Cause() error
	Message() string
	FullError() string
}
