package valval

import (
	"errors"
	"fmt"
	"gopkg.in/guregu/null.v2"
	"regexp"
	"unicode/utf8"
)

func errf(str string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(str, args...))
}

func GreaterThan(min float64) ValidatorFunc {
	return NewFloatValidator(func(val float64) error {
		if val <= min {
			return errf("must be greater than %g", val)
		}
		return nil
	})
}

func Min(min float64) ValidatorFunc {
	return NewFloatValidator(func(val float64) error {
		if val < min {
			return errf("must be %g and grater", val)
		}
		return nil
	})
}

func LessThan(max float64) ValidatorFunc {
	return NewFloatValidator(func(val float64) error {
		if val >= max {
			return errf("must be less than %g", val)
		}
		return nil
	})
}

func Max(max float64) ValidatorFunc {
	return NewFloatValidator(func(val float64) error {
		if val > max {
			return errf("must be %g or less", val)
		}
		return nil
	})
}

func Regexp(r *regexp.Regexp) ValidatorFunc {
	return NewStringValidator(func(val string) error {
		if !r.MatchString(val) {
			return errf("must be match to the regexp %s", r.String())
		}
		return nil
	})
}

func MinLength(min int) ValidatorFunc {
	return NewStringValidator(func(str string) error {
		if utf8.RuneCountInString(str) < min {
			return errf("length must be %d or greater", min)
		}
		return nil
	})
}

func MaxLength(max int) ValidatorFunc {
	return NewStringValidator(func(str string) error {
		if utf8.RuneCountInString(str) > max {
			return errf("length must be %d or less", max)
		}
		return nil
	})
}

func In(targetValues ...interface{}) ValidatorFunc {
	return func(val interface{}) error {
		for _, tv := range targetValues {
			if val == tv {
				return nil
			}
		}
		return errf("invalid value. allowed are %v", targetValues)
	}
}

func And(funcs ...ValidatorFunc) ValidatorFunc {
	return func(val interface{}) error {
		for _, f := range funcs {
			if err := f(val); err != nil {
				return err
			}
		}
		return nil
	}
}

func Or(funcs ...ValidatorFunc) ValidatorFunc {
	return func(val interface{}) error {
		var firstError error
		for _, f := range funcs {
			if err := f(val); err == nil {
				return nil
			} else {
				if firstError == nil {
					firstError = err
				}
			}
		}
		return firstError
	}
}

func MinSliceLength(min int) SliceValidatorFunc {
	return func(slice []interface{}) error {
		if len(slice) < min {
			return errf("length must be %d or greater", min)
		}
		return nil
	}
}

func MaxSliceLength(max int) SliceValidatorFunc {
	return func(slice []interface{}) error {
		if len(slice) > max {
			return errf("length must be %d or less", max)
		}
		return nil
	}
}

func RequiredFields(fs ...string) ObjectValidatorFunc {
	return func(content map[string]interface{}) error {
		for _, f := range fs {
			if content[f] == nil {
				return errf("field %s is required", f)
			}
		}
		return nil
	}
}

func NuullableMinLength(min int) ValidatorFunc {
	return NewNullStringValidator(func(str null.String) error {
		if utf8.RuneCountInString(str.String) < min {
			return errf("length must be %d or greater", min)
		}
		return nil
	})
}


func NuullableMaxLength(min int) ValidatorFunc {
	return NewNullStringValidator(func(str null.String) error {
		if utf8.RuneCountInString(str.String) > min {
			return errf("length must be %d or greater", min)
		}
		return nil
	})
}
