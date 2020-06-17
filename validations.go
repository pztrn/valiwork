// Valiwork - validations that works.
// Copyright (c) 2020 by Stanislav Nikitin <pztrn@pztrn.name>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
// Optional might be used for passing parameters to validators, where
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
