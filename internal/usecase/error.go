package usecase

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrForbidden           = errors.New("forbidden")
	ErrNotFound            = errors.New("not found")
	ErrConflict            = errors.New("conflict")
	ErrInternalServerError = errors.New("internal server error")
)
