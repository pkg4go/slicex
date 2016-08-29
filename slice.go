package slicex

import "reflect"

func Contains(list interface{}, item interface{}) bool {
	s := reflect.ValueOf(list)

	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return false
	}

	items := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		items[i] = s.Index(i).Interface()
	}

	for _, v := range items {
		if v == item {
			return true
		}
	}

	return false
}
