package apperrors

import (
	"errors"
)

var ErrNotFound = errors.New("not found")
var ErrForeignKey = errors.New("foreign key")
var ErrDuplicateKey = errors.New("duplicate key violation")
