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
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"go.dev.pztrn.name/valiwork/validators"
)

const (
	testString = " I am test $tring"
)

// nolint:paralleltest
func TestRegisterValidator(t *testing.T) {
	initializeValidatorsStorage()

	testCases := []struct {
		ValidatorFunc validators.ValidatorFunc
		ValidatorName string
		ShouldFail    bool
	}{
		{
			ValidatorName: "string_test_validator",
			ValidatorFunc: func(thing interface{}, optional ...interface{}) []interface{} {
				return nil
			},
			ShouldFail: false,
		},
		// This case is about registering same validator function again.
		{
			ValidatorName: "string_test_validator",
			ValidatorFunc: func(thing interface{}, optional ...interface{}) []interface{} {
				return nil
			},
			ShouldFail: true,
		},
	}

	for _, testCase := range testCases {
		err := RegisterValidator(testCase.ValidatorName, testCase.ValidatorFunc)

		if !testCase.ShouldFail {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func BenchmarkRegisterValidator(b *testing.B) {
	initializeValidatorsStorage()

	for i := 0; i < b.N; i++ {
		_ = RegisterValidator("string_test_validator_"+strconv.Itoa(i),
			func(thing interface{}, optional ...interface{}) []interface{} {
				return nil
			},
		)
	}
}

func BenchmarkRegisterValidatorAsync(b *testing.B) {
	initializeValidatorsStorage()

	var waiter sync.WaitGroup

	for i := 0; i < b.N; i++ {
		waiter.Add(1)

		go func() {
			_ = RegisterValidator("string_test_validator_"+strconv.Itoa(i),
				func(thing interface{}, optional ...interface{}) []interface{} {
					return nil
				},
			)

			waiter.Done()
		}()

		waiter.Wait()
	}
}

// nolint:paralleltest
func TestValidate(t *testing.T) {
	initializeValidatorsStorage()

	_ = RegisterValidator("string_test1", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.HasPrefix(stringToValidate, " ") {
			errs = append(errs, errors.New("string starts with whitespace, invalid"))
		}

		return errs
	})

	errs := Validate(testString, "string_test1", nil)
	require.NotNil(t, errs)
	require.Len(t, errs, 1)
}

func BenchmarkValidate(b *testing.B) {
	b.StopTimer()

	initializeValidatorsStorage()

	_ = RegisterValidator("string_test1", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.HasPrefix(stringToValidate, " ") {
			errs = append(errs, errors.New("string starts with whitespace, invalid"))
		}

		if strings.Contains(stringToValidate, "$") {
			errs = append(errs, errors.New("string contains dollar sign, invalid"))
		}

		return errs
	})

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Validate(testString, "string_test1")
	}
}

func BenchmarkValidateAsync(b *testing.B) {
	b.StopTimer()

	initializeValidatorsStorage()

	_ = RegisterValidator("string_test1", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.HasPrefix(stringToValidate, " ") {
			errs = append(errs, errors.New("string starts with whitespace, invalid"))
		}

		if strings.Contains(stringToValidate, "$") {
			errs = append(errs, errors.New("string starts with whitespace, invalid"))
		}

		return errs
	})

	b.StartTimer()

	var waiter sync.WaitGroup

	for i := 0; i < b.N; i++ {
		waiter.Add(1)

		go func() {
			_ = Validate(testString, "string_test1")

			waiter.Done()
		}()

		waiter.Wait()
	}
}

// nolint:paralleltest
func TestValidateMany(t *testing.T) {
	initializeValidatorsStorage()

	_ = RegisterValidator("string_test1", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.HasPrefix(stringToValidate, " ") {
			errs = append(errs, errors.New("string starts with whitespace, invalid"))
		}

		return errs
	})

	_ = RegisterValidator("string_test2", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.HasSuffix(stringToValidate, " ") {
			errs = append(errs, errors.New("string ends with whitespace, invalid"))
		}

		return errs
	})

	errs := ValidateMany(testString, []string{"string_test1", "string_test2"}, nil)
	require.NotNil(t, errs)
	require.Len(t, errs, 1)
}

func BenchmarkValidateMany(b *testing.B) {
	b.StopTimer()

	initializeValidatorsStorage()

	_ = RegisterValidator("string_test1", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.HasPrefix(stringToValidate, " ") {
			errs = append(errs, errors.New("string starts with whitespace, invalid"))
		}

		return errs
	})

	_ = RegisterValidator("string_test2", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.Contains(stringToValidate, "$") {
			errs = append(errs, errors.New("string contains dollar sign, invalid"))
		}

		return errs
	})

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = ValidateMany(testString, []string{"string_test1", "string_test2"}, nil)
	}
}

func BenchmarkValidateManyAsync(b *testing.B) {
	b.StopTimer()

	initializeValidatorsStorage()

	_ = RegisterValidator("string_test1", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.HasPrefix(stringToValidate, " ") {
			errs = append(errs, errors.New("string starts with whitespace, invalid"))
		}

		return errs
	})

	_ = RegisterValidator("string_test2", func(thing interface{}, optional ...interface{}) []interface{} {
		var errs []interface{}

		stringToValidate, ok := thing.(string)
		if !ok {
			errs = append(errs, errors.New("not a string"))

			return errs
		}

		if strings.Contains(stringToValidate, "$") {
			errs = append(errs, errors.New("string contains dollar sign, invalid"))
		}

		return errs
	})

	b.StartTimer()

	var waiter sync.WaitGroup

	for i := 0; i < b.N; i++ {
		waiter.Add(1)

		go func() {
			_ = ValidateMany(testString, []string{"string_test1", "string_test2"}, nil)

			waiter.Done()
		}()

		waiter.Wait()
	}
}

// nolint:paralleltest
func TestUnregisterValidator(t *testing.T) {
	initializeValidatorsStorage()

	testCases := []struct {
		ValidatorFunc validators.ValidatorFunc
		ValidatorName string
	}{
		{
			ValidatorName: "string_test_validator",
			ValidatorFunc: func(thing interface{}, optional ...interface{}) []interface{} {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		err := RegisterValidator(testCase.ValidatorName, testCase.ValidatorFunc)
		require.Nil(t, err)

		err1 := UnregisterValidator(testCase.ValidatorName)
		require.Nil(t, err1)
	}
}

// nolint:paralleltest
func TestUnregisterValidatorNotRegisteredValidator(t *testing.T) {
	initializeValidatorsStorage()

	err := UnregisterValidator("this is definitely not registered thing")
	require.NotNil(t, err)
}

func BenchmarkUnregisterValidator(b *testing.B) {
	b.StopTimer()

	initializeValidatorsStorage()

	for i := 0; i < b.N; i++ {
		_ = RegisterValidator("string_test_validator_"+strconv.Itoa(i),
			func(thing interface{}, optional ...interface{}) []interface{} {
				return nil
			},
		)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = UnregisterValidator("string_test_validator_" + strconv.Itoa(i))
	}
}

func BenchmarkUnregisterValidatorAsync(b *testing.B) {
	b.StopTimer()

	initializeValidatorsStorage()

	for i := 0; i < b.N; i++ {
		_ = RegisterValidator("string_test_validator_"+strconv.Itoa(i),
			func(thing interface{}, optional ...interface{}) []interface{} {
				return nil
			},
		)
	}

	var waiter sync.WaitGroup

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		waiter.Add(1)

		go func() {
			_ = UnregisterValidator("string_test_validator_" + strconv.Itoa(i))

			waiter.Done()
		}()

		waiter.Wait()
	}
}
