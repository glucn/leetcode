package solution

import "strconv"

// Runtime 8 ms, faster than 89.77%
// Memory 6.6 MB, less than 50%

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	cache, write := 0, 0

	for i := range chars {
		if i == len(chars)-1 || chars[i] != chars[i+1] {
			chars[write] = chars[cache]
			write++
			if i > cache {
				countStr := strconv.Itoa(i - cache + 1)
				for j := range countStr {
					chars[write] = countStr[j]
					write++
				}
			}
			cache = i + 1
		}
	}

	return write
}
