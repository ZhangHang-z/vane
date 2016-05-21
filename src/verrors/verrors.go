package verrors

import (
	"errors"
)

var (
	ERR_HTTP_NOT_FOUND  = errors.New("Package Not Found")
	ERR_PTR_HELP_STRING = errors.New("This error for print some help string, usually omit it.")
)
