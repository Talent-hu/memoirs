package utils

import "reflect"

// RemoveSliceMap 结构体列表去重
func RemoveSliceMap(target []interface{})(result []interface{}) {
	n := len(target)
	for i := 0; i < n; i++ {
		state := false
		for j := i+1; j < n ; j++ {
			if reflect.DeepEqual(target[i],target[j]) {
				state = true
				break
			}
		}
		if !state {
			result = append(result,target[i])
		}
	}
	return
}
