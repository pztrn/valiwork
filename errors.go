package valiwork

import (
	// stdlib
	"errors"
)

var (
	// ErrValidatorAlreadyRegistered appears when validator's name
	// passed to RegisterValidator function already used for other
	// validator function.
	ErrValidatorAlreadyRegistered = errors.New("validator with such name already registered")

	// ErrValidatorNotRegistered appears when trying to unregister
	// not registered validator function.
	ErrValidatorNotRegistered = errors.New("validator with such name wasn't registered")
)
