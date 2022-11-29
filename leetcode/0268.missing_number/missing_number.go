package leetcode

func missingNumber(nums []int) int {
	n := len(nums)
	i := 0
	for i < len(nums) {
		j := nums[i]
		if nums[i] < n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
			continue
		}
	}

	for idx, num := range nums {
		if idx != num {
			return idx
		}
	}
	return n
}
