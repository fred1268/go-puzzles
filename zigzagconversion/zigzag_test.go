package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	t.Parallel()
	testData := []struct {
		name     string
		str      string
		rows     int
		expected string
	}{
		{
			name:     "PAYPALISHIRING-3",
			str:      "PAYPALISHIRING",
			rows:     3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "PAYPALISHIRING-4",
			str:      "PAYPALISHIRING",
			rows:     4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "A-1",
			str:      "A",
			rows:     1,
			expected: "A",
		},
	}
	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := convert(tt.str, tt.rows)
			if tt.expected != result {
				t.Errorf("wanted: '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
