package validators

// ValidatorFunc represents signature for data validation function.
type ValidatorFunc func(thing interface{}, optional ...interface{}) []interface{}
