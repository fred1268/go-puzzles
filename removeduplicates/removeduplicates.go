package removeduplicates

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	removed := 0
	last := len(nums)
	for i := 0; i < last; i++ {
		count := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				break
			}
			count++
		}
		if count > 2 {
			for j := i + 2; j < last-count+2; j++ {
				nums[j] = nums[j+count-2]
			}
			for j := last - count + 2; j < last; j++ {
				nums[j] = -1
			}
			removed += count - 2
			last -= count - 2
			i = i + 1
		}
	}
	return len(nums) - removed
}
