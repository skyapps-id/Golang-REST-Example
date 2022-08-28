package internal_const

import (
	"golang-rest-example/utils/errors"
	"net/http"
)

var (
	//ErrAuthPrixaRequestInvalidSourceHeader is
	ErrAuthPrixaRequestInvalidSourceHeader = func() error {
		return errors.ErrBase.New("invalid header x-source").WithProperty(errors.ErrCodeProperty, 2201).WithProperty(errors.ErrHttpCodeProperty, http.StatusUnauthorized)
	}

	//ErrAuthPrixaRequestInvalidTimestampHeader is
	ErrAuthPrixaRequestInvalidTimestampHeader = func() error {
		return errors.ErrBase.New("invalid header x-timestamp").WithProperty(errors.ErrCodeProperty, 2202).WithProperty(errors.ErrHttpCodeProperty, http.StatusUnauthorized)
	}

	//ErrAuthPrixaRequestExpired is
	ErrAuthPrixaRequestExpired = func() error {
		return errors.ErrBase.New("request is expired").WithProperty(errors.ErrCodeProperty, 2203).WithProperty(errors.ErrHttpCodeProperty, http.StatusUnauthorized)
	}

	//ErrAuthPrixaRequestInvalidSignature is
	ErrAuthPrixaRequestInvalidSignature = func() error {
		return errors.ErrBase.New("invalid signature").WithProperty(errors.ErrCodeProperty, 2203).WithProperty(errors.ErrHttpCodeProperty, http.StatusUnauthorized)
	}

	//ErrAuthPrixaRequestData is
	ErrAuthPrixaRequestData = func() error {
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
