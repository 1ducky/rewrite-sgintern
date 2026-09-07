package mapper

import (
	"errors"
)

type ErrorMapping struct {
	NotFound        error
	Duplicate       error
	RelatedNotFound error
	StillReferenced error
	RequiredField   error
	ValueTooLong    error
	Retryable       error
	Connection      error
}

var (
	ErrNoRows          = errors.New("Row not found")
	ErrDuplicate       = errors.New("Duplicate record")
	ErrRelatedNotFound = errors.New("Related record not found")
	ErrStillReferenced = errors.New("Still referenced")
	ErrRequiredField   = errors.New("Required field is empty")
	ErrValueTooLong    = errors.New("Value is too long")
	ErrRetryable       = errors.New("Retryable error")
	ErrConnection      = errors.New("Connection error")
)

func GenerateDefaultMapping(mapp ErrorMapping) ErrorMapping {
	if mapp.NotFound == nil {
		mapp.NotFound = ErrNoRows
	}
	if mapp.Duplicate == nil {
		mapp.Duplicate = ErrDuplicate
	}
	if mapp.RelatedNotFound == nil {
		mapp.RelatedNotFound = ErrRelatedNotFound
	}
	if mapp.StillReferenced == nil {
		mapp.StillReferenced = ErrStillReferenced
	}
	if mapp.RequiredField == nil {
		mapp.RequiredField = ErrRequiredField
	}
	if mapp.ValueTooLong == nil {
		mapp.ValueTooLong = ErrValueTooLong
	}
	if mapp.Retryable == nil {
		mapp.Retryable = ErrRetryable
	}
	if mapp.Connection == nil {
		mapp.Connection = ErrConnection
	}
	return mapp
}
