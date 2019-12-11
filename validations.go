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

// Validate launches validation function and returns it's result to
// caller. Optional might be used for passing additional options to
// validators.
func Validate(thing interface{}, validatorName string, optional ...interface{}) []interface{} {
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

// ValidateMany launches many validators using one-line-call.
// Optional might be used for passing parameters to validators, wher
// key is a validator name and value (which is []interface{})
// is a slice of parameters.
func ValidateMany(thing interface{}, validatorNames []string, optional map[string][]interface{}) []interface{} {
	var errs []interface{}

	for _, validator := range validatorNames {
		validatorParams, found := optional[validator]
		if !found {
			validatorParams = make([]interface{}, 0)
		}

		errs1 := Validate(thing, validator, validatorParams...)
		if len(errs1) > 0 {
			errs = append(errs, errs1...)
		}
	}

	return errs
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
