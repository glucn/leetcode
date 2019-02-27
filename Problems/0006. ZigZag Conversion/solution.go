package solution

// Runtime 12 ms
// Memory 5.9 MB

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	cache := make([][]byte, numRows)

	zigzag := numRows*2 - 2

	for i := 0; i < len(s); i++ {
		row := i % zigzag

		if row > numRows-1 {
			row = zigzag - row
		}
		cache[row] = append(cache[row], s[i])
	}

	var result []byte
	for i := 0; i < numRows; i++ {
		result = append(result, cache[i]...)
	}

	return string(result[:])
}
