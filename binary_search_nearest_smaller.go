package main

func BinarySearchNearestSmaller(arr []int, target int) (int, bool) {
	lo := 0
	hi := len(arr) - 1
	result_idx := -1

	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] < target {
			result_idx = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	if result_idx != 0 {
		return arr[result_idx], true
	} else {
		return -1, false
	}
}
