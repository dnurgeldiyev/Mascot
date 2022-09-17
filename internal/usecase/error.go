package usecase

import "errors"

var (
	ErrBadRequest          = errors.New("ErrBadRequest")
	ErrNotFound            = errors.New("ErrNotFound")
	ErrConflict            = errors.New("ErrConflict")
	ErrInternalServerError = errors.New("ErrInternalServerError")
	ErrNotEnoughMoneyCode  = errors.New("ErrNotEnoughMoneyCode")
)
