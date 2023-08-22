package utils

import (
	"fmt"
	"reflect"
	"sort"
)

func InOfStr(str string, arrayList []string) bool {
	for _, v := range arrayList {
		if str == v {
			return true
		}
	}
	return false
}

func InOfInt(obj int, arrayList []int) bool {
	for _, v := range arrayList {
		if obj == v {
			return true
		}
	}
	return false
}

func InOfInt32(obj int32, arrayList []int32) bool {
	for _, v := range arrayList {
		if obj == v {
			return true
		}
	}
	return false
}

func InOfInt64(obj int64, arrayList []int64) bool {
	for _, v := range arrayList {
		if obj == v {
			return true
		}
	}
	return false
}

func InOfType(obj interface{}, arrayList []interface{}, dataType string) bool {
	for _, v := range arrayList {
		if obj == v {
			return true
		}
	}
	return false
}

func InOfT[T string | int | int64 | int32](t T, arrayT []T) bool {
	for _, v := range arrayT {
		if t == v {
			return true
		}
	}
	return false
}

func InWithSlice[T string | int | int64](source []T, target []T) bool {
	for _, s := range source {
		if InOfT(s, target) {
			return true
		}
	}
	return false
}

//func In(obj interface{}, arrayList []interface{}) bool {
//	for _,v:=range arrayList{
//		fmt.Println(v)
//		if obj == v{
//			return true
//		}
//	}
//	return false
//}

// In example:
// r, err := In([]interface{}{1, "two", 3}, "two")
func In(haystack interface{}, needle interface{}) (bool, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == needle {
				return true, nil
			}
		}
		return false, nil
	}
	return false, fmt.Errorf("UnSupport haystack kind: %s", kind)
}

// 二分查找
func SortInIntSlice(haystack []int, needle int) bool {
	sort.Ints(haystack)
	index := sort.SearchInts(haystack, needle)
	return index < len(haystack) && haystack[index] == needle
}

func InIntSliceMapKeyFunc(haystack []int) func(int) bool {
	set := make(map[int]struct{})

	for _, e := range haystack {
		set[e] = struct{}{}
	}

	return func(needle int) bool {
		_, ok := set[needle]
		return ok
	}
}

func Pop(haystack interface{}, needle interface{}) (bool, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == needle {
				return true, nil
			}
		}
		return false, nil
	}
	return false, fmt.Errorf("UnSupport haystack kind: %s", kind)
}
