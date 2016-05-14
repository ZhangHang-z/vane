package parser

import (
	"errors"
)

var (
	ERR_RC_FILE_NOT_FOUND   = errors.New(".vanerc file not found.")
	ERR_JSON_FILE_NOT_FOUND = errors.New("vane.json file not found.")
	ERR_RC_FILE_CONF        = errors.New(".vanerc file configuration error.")
	ERR_JSON_FILE_CONF      = errors.New("vane.json file configuration error.")
)
