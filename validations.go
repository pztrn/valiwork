package valiwork

import (
	// stdlib
	"log"

	// local
	"go.dev.pztrn.name/valiwork/validators"
)

var (
	registeredValidators map[string]validators.ValidatorFunc
)

// nolint
func init() {
	initializeValidatorsStorage()
}

func initializeValidatorsStorage() {
	registeredValidators = make(map[string]validators.ValidatorFunc)
}

// RegisterValidator registers validation function for later calling.
func RegisterValidator(validatorName string, validator validators.ValidatorFunc) error {
	if DEBUG {
		log.Println("Trying to register validator: '" + validatorName + "'...")
	}

	_, found := registeredValidators[validatorName]

	if found {
		if DEBUG {
			log.Println("Validator already registered!")
		}

		return ErrValidatorAlreadyRegistered
	}

	registeredValidators[validatorName] = validator

	return nil
}

// UnregisterValidator removes registered validator from list of known
// validators.
func UnregisterValidator(validatorName string) error {
	if DEBUG {
		log.Println("Trying to unregister validator '" + validatorName + "'...")
	}

	_, found := registeredValidators[validatorName]

	if !found {
		if DEBUG {
			log.Println("Validator wasn't registered!")
		}

		return ErrValidatorNotRegistered
	}

	delete(registeredValidators, validatorName)

	return nil
}

// Validate launches validation function and returns it's result to
// caller.
func Validate(validatorName string, thing interface{}, optional ...interface{}) []interface{} {
	var errs []interface{}

	validator, found := registeredValidators[validatorName]

	if !found {
		errs = append(errs, ErrValidatorNotRegistered)
		return errs
	}

	errs1 := validator(thing, optional...)
	if len(errs1) > 0 {
		errs = append(errs, errs1...)
	}

	return errs
}
