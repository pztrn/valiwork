package main

import (
	"strings"
	// stdlib
	"errors"
	"log"

	// other
	"go.dev.pztrn.name/valiwork"
)

const (
	stringValidatorName = "string_validate_things"
)

func main() {
	log.Println("Starting validation example...")
	log.Println("WARN: to see additional valiwork output define 'VALIWORK_DEBUG' environment variable and set it to 'true'!")

	//stringToValidate := " I am pretty b@d $tring"

	_ = valiwork.RegisterValidator(stringValidatorName, stringValidator)
}

func stringValidator(thing interface{}, optional ...interface{}) []interface{} {
	var errs []interface{}

	stringToValidate, ok := thing.(string)
	if !ok {
		errs = append(errs, errors.New("passed value is not a string"))
		return errs
	}

	// Are string begins with spaces?
	if strings.HasPrefix(stringToValidate, " ") {
		errs = append(errs, errors.New("string begins with space"))
	}

	// Does string contains any special characters?

	return errs
}
