package mergearrays

import "fmt"

func mergeArrays(array1, array2 []int) ([]int, []int) {
	for i := 0; i < len(array1); i++ {
		fmt.Printf("%v, %v\n", array1, array2)
		for j := 0; j < len(array2); j++ {
			if array2[i] < array1[i] {
				array1[i], array2[j] = array2[j], array1[i]
				for {
					if j+1 > len(array2) || array2[j] < array2[j+1] {
						break
					}
					if array2[j] > array2[j+1] {
						array2[j], array2[j+1] = array2[j+1], array2[j]
					}
					j++
				}
				break
			}
		}
	}
	return array1, array2
}
