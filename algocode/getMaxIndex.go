package main

import "fmt"

func GetMaxIndex(values []*int) int {
	if len(values) == 0 || values == nil {
		return 0
	}
	max_index := 0
	var max_value *int

	for index, val := range values {
		fmt.Println(index, *val)
		if val == nil {
			continue
		}
		if max_value == nil || *val > *max_value {
			max_value = val
			max_index = index
		}
	}

	return max_index
}

//a := 3
//b := 9
//c := 5
//d := 9
//arr := []*int{&a, &b, &c, &d}
//fmt.Println(GetMaxIndex(arr))
