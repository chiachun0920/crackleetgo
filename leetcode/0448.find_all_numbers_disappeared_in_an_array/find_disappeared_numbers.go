package leetcode

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}

	result := []int{}
	for i, num := range nums {
		if i+1 != num {
			result = append(result, i+1)
		}
	}
	return result
}
