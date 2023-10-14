package longestsubstring

import "testing"

func TestLongestSubstring(t *testing.T) {
	testData := []struct {
		name     string
		str      string
		expected string
	}{
		{
			name:     "abcabcbb",
			str:      "abcabcbb",
			expected: "abc",
		},
		{
			name:     "bbbbb",
			str:      "bbbbb",
			expected: "b",
		},
		{
			name:     "pwwkew",
			str:      "pwwkew",
			expected: "wke",
		},
		{
			name:     "pwwkewkep",
			str:      "pwwkewkep",
			expected: "wkep",
		},
		{
			name:     "ABCDEFGABEF",
			str:      "ABCDEFGABEF",
			expected: "ABCDEFG",
		},
	}
	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			_, result := longestSubstring(tt.str)
			if tt.expected != result {
				t.Errorf("wanted: '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
