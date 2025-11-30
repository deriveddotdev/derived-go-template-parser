package utils

import (
	"reflect"
)

func SliceHas(slice interface{}, key string, value interface{}) bool {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		return false
	}

	targetVal := reflect.ValueOf(value)

	// Iterate through the slice
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)

		// Handle different types of elements
		switch elem.Kind() {
		case reflect.Struct:
			// Get the field by name
			field := elem.FieldByName(key)
			if !field.IsValid() {
				continue
			}
			// Compare the values
			if reflect.DeepEqual(field.Interface(), targetVal.Interface()) {
				return true
			}

		case reflect.Map:
			// Get the map value by key
			mapVal := elem.MapIndex(reflect.ValueOf(key))
			if !mapVal.IsValid() {
				continue
			}
			// Compare the values
			if reflect.DeepEqual(mapVal.Interface(), targetVal.Interface()) {
				return true
			}

		case reflect.Interface, reflect.Ptr:
			// Dereference if necessary and recurse
			if !elem.IsNil() {
				if SliceHas(elem.Elem().Interface(), key, value) {
					return true
				}
			}
		}
	}

	return false
}
