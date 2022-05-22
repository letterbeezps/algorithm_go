package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	if nums[0] > nums[len(nums)-1] {
		l, r := 0, len(nums)-1
		for l < r {
			mid := (l + r + 1) / 2
			if nums[mid] >= nums[0] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if target >= nums[0] {
			left = 0
			right = l
		} else {
			left = l + 1
			right = len(nums) - 1
		}
	}

	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if target == nums[left] {
		return left
	} else {
		return -1
	}
}
