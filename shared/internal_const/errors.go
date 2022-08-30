package internal_const

import (
	"golang-rest-example/utils/errors"
	"net/http"
)

var (
	//ErrAuthRequestInvalidSourceHeader is
	ErrAuthRequestInvalidSourceHeader = func() error {
		return errors.ErrBase.New("invalid header x-source").WithProperty(errors.ErrCodeProperty, 400).WithProperty(errors.ErrHttpCodeProperty, http.StatusUnauthorized)
	}

	//ErrAuthRequestInvalidTimestampHeader is
	ErrAuthRequestInvalidTimestampHeader = func() error {
		return errors.ErrBase.New("invalid header x-timestamp").WithProperty(errors.ErrCodeProperty, 400).WithProperty(errors.ErrHttpCodeProperty, http.StatusUnauthorized)
	}

	//ErrAuthRequestExpired is
	ErrAuthRequestExpired = func() error {
		return errors.ErrBase.New("request is expired").WithProperty(errors.ErrCodeProperty, 401).WithProperty(errors.ErrHttpCodeProperty, http.StatusUnauthorized)
	}

	//ErrAuthRequestInvalidSignature is
	ErrAuthRequestInvalidSignature = func() error {
		return errors.ErrBase.New("invalid signature").WithProperty(errors.ErrCodeProperty, 401).WithProperty(errors.ErrHttpCodeProperty, http.StatusUnauthorized)
	}

	//ErrAuthRequestData is
	ErrAuthRequestData = func() error {
		return errors.ErrBase.New("user not have permission").WithProperty(errors.ErrCodeProperty, http.StatusForbidden).WithProperty(errors.ErrHttpCodeProperty, http.StatusForbidden)
	}

	//ErrRecordNotFound is
	ErrRecordNotFound = func() error {
		return errors.ErrBase.New("record not found").WithProperty(errors.ErrCodeProperty, http.StatusNotFound).WithProperty(errors.ErrHttpCodeProperty, http.StatusNotFound)
	}

	// ErrBadRequest
	ErrBadRequest = func(err error) error {
		return errors.ErrBase.New(err.Error()).WithProperty(errors.ErrCodeProperty, http.StatusBadRequest).WithProperty(errors.ErrHttpCodeProperty, http.StatusBadRequest)
	}

	// ErrBadRequest
	ErrBadGateway = func(err error) error {
		return errors.ErrBase.New(err.Error()).WithProperty(errors.ErrCodeProperty, http.StatusBadGateway).WithProperty(errors.ErrHttpCodeProperty, http.StatusBadGateway)
	}
)
