package isnil

import (
	"reflect"
)

// IsNil returns true if a given value is nil.
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Ptr && v.IsNil()
}
