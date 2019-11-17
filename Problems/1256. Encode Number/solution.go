package solution

import "strconv"

// Runtime 0 ms, faster than ?%
// Memory 2 MB, less than ?%

func encode(num int) string {
	if num == 0 {
		return ""
	}
	s := strconv.FormatInt(int64(num+1), 2)
	return s[1:]
}
