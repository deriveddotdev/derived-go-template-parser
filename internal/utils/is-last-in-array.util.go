package utils

func IsLast(index int, slice interface{}) bool {
	s, ok := slice.([]interface{}) // Type assert to slice
	if !ok {
		return false // Not a slice, return false
	}
	return index == len(s)-1
}
