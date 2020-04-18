package validator

import "fmt"

func (v *Validator) EqualToField(field string, value interface{}, toEqualField string, toEqualValue interface{}) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if value != toEqualField {
		v.Errors[field] = fmt.Sprintf("%s must be equal to %s", field, toEqualField)
		return false
	}
	return true
}
