package mergearrays

import "testing"

func equals(array1, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

func TestMergeArray(t *testing.T) {
	testData := []struct {
		name      string
		array1    []int
		array2    []int
		expected1 []int
		expected2 []int
	}{
		{
			name:      "4-6",
			array1:    []int{1, 3, 5, 7},
			array2:    []int{0, 2, 4, 6, 8, 9},
			expected1: []int{0, 1, 2, 3},
			expected2: []int{4, 5, 6, 7, 8, 9},
		},
		{
			name:      "2+3",
			array1:    []int{10, 12},
			array2:    []int{5, 18, 20},
			expected1: []int{5, 10},
			expected2: []int{12, 18, 20},
		},
		{
			name:      "2+3",
			array1:    []int{1, 3, 4, 7, 8, 9},
			array2:    []int{2, 5, 10},
			expected1: []int{1, 2, 3, 4, 5, 7},
			expected2: []int{8, 9, 10},
		},
	}
	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			result1, result2 := mergeArrays(tt.array1, tt.array2)
			if !equals(tt.expected1, result1) || !equals(tt.expected2, result2) {
				t.Errorf("wanted: '%v', '%v', got '%v', '%v'", tt.expected1, tt.expected2, result1, result2)
			}
		})
	}
}
