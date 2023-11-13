package palindromicsubstring

import "testing"

func TestPalintromicSubstring(t *testing.T) {
	testData := []struct {
		name     string
		expected string
	}{
		{
			name:     "babad",
			expected: "bab",
		},
		{
			name:     "cbbd",
			expected: "bb",
		},
		{
			name:     "aracecar",
			expected: "racecar",
		},
		{
			name:     "deifieddhsjdh",
			expected: "deified",
		},
		{
			name:     "jkljkdeifiedopop",
			expected: "deified",
		},
		{
			name:     "deifiedoiop",
			expected: "deified",
		},
		{
			// Sir, I demand, I am a maid named Iris
			name:     "siridemandiamamaidnamediris",
			expected: "siridemandiamamaidnamediris",
		},
	}
	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := palindromicsubstring(tt.name)
			if tt.expected != result {
				t.Errorf("wanted: '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
