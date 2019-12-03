package valiwork

import (
	// stdlib
	"log"
	"sync"

	// local
	"go.dev.pztrn.name/valiwork/validators"
)

var (
	registeredValidators map[string]validators.ValidatorFunc
	rvMutex              sync.RWMutex
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

	//rvMutex.RLock()
	_, found := registeredValidators[validatorName]
	//rvMutex.RUnlock()

	if found {
		if DEBUG {
			log.Println("Validator already registered!")
		}

		return ErrValidatorAlreadyRegistered
	}

	//rvMutex.Lock()
	registeredValidators[validatorName] = validator
	//rvMutex.Unlock()

	return nil
}

// UnregisterValidator removes registered validator from list of known
// validators.
func UnregisterValidator(validatorName string) error {
	if DEBUG {
		log.Println("Trying to unregister validator '" + validatorName + "'...")
	}

	//rvMutex.RLock()
	_, found := registeredValidators[validatorName]
	//rvMutex.RUnlock()

	if !found {
		if DEBUG {
			log.Println("Validator wasn't registered!")
		}

		return ErrValidatorNotRegistered
	}

	//rvMutex.Lock()
	delete(registeredValidators, validatorName)
	//rvMutex.Unlock()

	return nil
}

// Validate launches validation function and returns it's result to
// caller.
func Validate(validatorName string, thing interface{}, optional ...interface{}) []interface{} {
	var errs []interface{}

	//rvMutex.RLock()
	validator, found := registeredValidators[validatorName]
	//rvMutex.RUnlock()

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
