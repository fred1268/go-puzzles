package zigzagconversion

import (
	"fmt"
)

func convert(s string, rows int) string {
	lines := make([]string, rows)
	row, dir := 0, 1
	for _, c := range s {
		lines[row] = fmt.Sprintf("%s%c", lines[row], c)
		row = row + dir
		if row == 0 || row == rows-1 {
			dir = -dir
		}
	}
	result := ""
	for _, str := range lines {
		result = fmt.Sprintf("%s%s", result, str)
	}
	return result
}
