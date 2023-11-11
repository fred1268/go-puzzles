package removeduplicates

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

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()
	testData := []struct {
		name     string
		array    []int
		expected []int
	}{
		{
			name:     "111223",
			array:    []int{1, 1, 1, 2, 2, 3},
			expected: []int{1, 1, 2, 2, 3, -1},
		},
		{
			name:     "001111233",
			array:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected: []int{0, 0, 1, 1, 2, 3, 3, -1, -1},
		},
		{
			name:     "00111123333",
			array:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 3, 3},
			expected: []int{0, 0, 1, 1, 2, 3, 3, -1, -1, -1, -1},
		},
	}
	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, _ := removeDuplicates(tt.array)
			if !equals(tt.expected, result) {
				t.Errorf("wanted: '%v', got '%v'", tt.expected, result)
			}
		})
	}
}
