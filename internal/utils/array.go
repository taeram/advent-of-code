package utils

import "reflect"

func InArray(needle interface{}, haystack interface{}) bool {
	hasNeedle := false
	if reflect.TypeOf(haystack).Kind() == reflect.Slice {
		s := reflect.ValueOf(haystack)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				hasNeedle = true
				break
			}
		}
	}

	return hasNeedle
}
