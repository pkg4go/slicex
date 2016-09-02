package slicex

import "reflect"

func Contains(list interface{}, item interface{}) bool {
	s := reflect.ValueOf(list)

	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < s.Len(); i++ {
		it := s.Index(i).Interface()
		if it == item {
			return true
		}
	}

	return false
}

func Index(list interface{}, item interface{}) int {
	s := reflect.ValueOf(list)

	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return -1
	}

	for i := 0; i < s.Len(); i++ {
		it := s.Index(i).Interface()
		if it == item {
			return i
		}
	}

	return -1
}
