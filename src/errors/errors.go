package errors

import (
	stdErr "errors"
)

var (
	ERR_HTTP_NOT_FOUND      = stdErr.New("Package Not Found")
	ERR_RC_FILE_NOT_FOUND   = stdErr.New(".vanerc file not found.")
	ERR_JSON_FILE_NOT_FOUND = stdErr.New("vane.json file not found.")
	ERR_RC_FILE_CONF        = stdErr.New(".vanerc file configuration error.")
	ERR_JSON_FILE_CONF      = stdErr.New("vane.json file configuration error.")
	ERR_PTR_HELP_STRING     = stdErr.New("This error for print some help string, usually omit it.")
)
