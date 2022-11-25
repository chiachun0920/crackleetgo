package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	currentReplaceIdx := 0

	for _, value := range nums {
		if value == nums[currentReplaceIdx] {
			continue
		}

		currentReplaceIdx++

		nums[currentReplaceIdx] = value
	}

	return currentReplaceIdx + 1
}
