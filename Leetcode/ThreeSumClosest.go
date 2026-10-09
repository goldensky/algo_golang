package main

import (
	"math"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest_sum := nums[0] + nums[1] + nums[2]
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		min_sum := nums[i] + nums[i+1] + nums[i+2]
		if min_sum > target {
			if abs(min_sum-target) < abs(closest_sum-target) {
				closest_sum = min_sum
				break
			}
		}

		max_sum := nums[i] + nums[n-2] + nums[n-1]
		if max_sum < target {
			if abs(max_sum-target) < abs(closest_sum-target) {
				closest_sum = max_sum
				continue
			}
		}

		lo, hi := i+1, n-1
		for lo < hi {
			current_sum := nums[i] + nums[lo] + nums[hi]
			if abs(current_sum-target) < abs(closest_sum-target) {
				closest_sum = current_sum
			}

			if current_sum < target {
				lo++
			} else if current_sum > target {
				hi--
			} else {
				return target
			}
		}
	}

	return closest_sum
}

func threeSumClosest_1(nums []int, target int) int {
	sort.Ints(nums)
	closest_sum := nums[0] + nums[1] + nums[2]
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		min_sum := nums[i] + nums[i+1] + nums[i+2]
		if min_sum > target {
			if math.Abs(float64(min_sum-target)) < math.Abs(float64(closest_sum-target)) {
				closest_sum = min_sum
				break
			}
		}

		max_sum := nums[i] + nums[n-2] + nums[n-1]
		if max_sum < target {
			if math.Abs(float64(max_sum-target)) < math.Abs(float64(closest_sum-target)) {
				closest_sum = max_sum
				continue
			}
		}

		lo, hi := i+1, n-1
		for lo < hi {
			current_sum := nums[i] + nums[lo] + nums[hi]
			if math.Abs(float64(current_sum-target)) < math.Abs(float64(closest_sum-target)) {
				closest_sum = current_sum
			}

			if current_sum < target {
				lo++
			} else if current_sum > target {
				hi--
			} else {
				return target
			}
		}
	}

	return closest_sum
}
