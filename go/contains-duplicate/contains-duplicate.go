package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	l := len(nums)
	for i := 0; i < l; i++ {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}

// passed
