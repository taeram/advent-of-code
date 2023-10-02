package utils

import "reflect"

func InArray(needle interface{}, haystack interface{}) bool {
	hasNeedle := false
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
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
