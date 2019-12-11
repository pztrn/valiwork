package valiwork

import (
	// stdlib
	"errors"
	"strconv"
	"strings"
	"sync"
	"testing"

	// local
	"go.dev.pztrn.name/valiwork/validators"

	// other
	"github.com/stretchr/testify/require"
)

func TestRegisterValidator(t *testing.T) {
	initializeValidatorsStorage()

	testCases := []struct {
		ValidatorName string
		ValidatorFunc validators.ValidatorFunc
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

	var w sync.WaitGroup

	for i := 0; i < b.N; i++ {
		w.Add(1)

		go func() {
			_ = RegisterValidator("string_test_validator_"+strconv.Itoa(i),
				func(thing interface{}, optional ...interface{}) []interface{} {
					return nil
				},
			)
			w.Done()
		}()

		w.Wait()
	}
}

func TestValidate(t *testing.T) {
	initializeValidatorsStorage()

	testString := " I am test string"

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

	testString := " I am test $tring"

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

	testString := " I am test $tring"

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

	var w sync.WaitGroup

	for i := 0; i < b.N; i++ {
		w.Add(1)

		go func() {
			_ = Validate(testString, "string_test1")

			w.Done()
		}()

		w.Wait()
	}
}

func TestValidateMany(t *testing.T) {
	initializeValidatorsStorage()

	testString := " I am test string"

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

	testString := " I am test $tring"

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

	testString := " I am test $tring"

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

	var w sync.WaitGroup

	for i := 0; i < b.N; i++ {
		w.Add(1)

		go func() {
			_ = ValidateMany(testString, []string{"string_test1", "string_test2"}, nil)

			w.Done()
		}()

		w.Wait()
	}
}

func TestUnregisterValidator(t *testing.T) {
	initializeValidatorsStorage()

	testCases := []struct {
		ValidatorName string
		ValidatorFunc validators.ValidatorFunc
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

	var w sync.WaitGroup

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		w.Add(1)

		go func() {
			_ = UnregisterValidator("string_test_validator_" + strconv.Itoa(i))
			w.Done()
		}()

		w.Wait()
	}
}
