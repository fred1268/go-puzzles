package palindromicsubstring

func palindromicsubstring(input string) string {
	if len(input) == 1 {
		return input
	}
	if len(input) == 2 {
		if input[0] == input[1] {
			return input
		}
		return input[1:]
	}
	maxLength := -1
	position := -1
	for i := 2; i < len(input); i++ {
		length := 0
		for {
			if input[i+length] != input[i-2-length] {
				break
			}
			length++
			if i+length >= len(input) || i-2-length < 0 {
				break
			}
		}
		if 2*length+1 > maxLength {
			maxLength = 2*length + 1
			position = i - 2 - length + 1
		}
		length = 0
		for {
			if input[i+length] != input[i-1-length] {
				break
			}
			length++
			if i+length >= len(input) || i-2-length < 0 {
				break
			}
		}
		if 2*length > maxLength {
			maxLength = 2 * length
			position = i - 1 - length + 1
		}
	}
	return input[position : position+maxLength]
}
