package validator

import "fmt"

func (v *Validator) MinLength(field, value string, minLength int) bool  {
	if _, ok := v.Errors[field]; ok {
		return false
	}
	if len(value) < minLength {
		v.Errors[field] = fmt.Sprintf("%s must atleast be (%d) long", field, minLength)
		return false
	}
	return true
}
