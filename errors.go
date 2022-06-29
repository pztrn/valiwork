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
