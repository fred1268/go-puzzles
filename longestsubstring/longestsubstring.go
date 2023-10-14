package longestsubstring

func longestSubstring(str string) (int, string) {
	var pos int
	var substring string
	uniq := make(map[byte]struct{})
	for start := 0; start < len(str)-len(substring); start++ {
		for i := start; i < len(str); i++ {
			c := str[i]
			if _, ok := uniq[c]; ok {
				if len(substring) < len(str[start:i]) {
					substring = str[start:i]
					pos = start
				}
				clear(uniq)
				break
			}
			uniq[c] = struct{}{}
		}
		if len(uniq) != 0 && len(substring) < len(str[start:]) {
			substring = str[start:]
			pos = start
			break
		}
	}
	return pos, substring
}
