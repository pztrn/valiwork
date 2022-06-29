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

package main

import (
	"errors"
	"log"
	"strings"

	"go.dev.pztrn.name/valiwork"
)

const (
	stringValidatorName = "string_validate_things"
)

func main() {
	log.Println("Starting validation example...")
	log.Println("WARN: to see additional valiwork output define 'VALIWORK_DEBUG' environment variable and set it to 'true'!")

	// stringToValidate := " I am pretty b@d $tring"

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
