package removeduplicates

import "fmt"

func removeDuplicates(nums []int) ([]int, int) {
	if len(nums) < 2 {
		return nums, len(nums)
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
		fmt.Printf("Found %d times %d\n", count, nums[i])
		if count > 2 {
			for j := i; j < last; j++ {
				switch {
				case j-i < 2:
					fmt.Printf("skipping\n")
				case j-i >= 2 && j < last-count+2:
					nums[j] = nums[j+count-2]
					fmt.Printf("(shifting) nums: %v\n", nums)
				case j >= last-count+2:
					nums[j] = -1
					fmt.Printf("(setting -1) nums: %v\n", nums)
				}
			}
			removed += count - 2
			last -= count - 2
			i = i + 1
		}
	}
	return nums, len(nums) - removed
}
