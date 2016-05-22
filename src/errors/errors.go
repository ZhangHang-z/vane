package errors

type vaneError struct {
	msg string
}

func New(msg string) *vaneError {
	return &vaneError{msg}
}

func (err *vaneError) Error() string {
	return err.msg
}

var (
	ERR_HTTP_NOT_FOUND  = New("Package Not Found")
	ERR_PTR_HELP_STRING = New("This error for print some help string, usually omit it.")
)
