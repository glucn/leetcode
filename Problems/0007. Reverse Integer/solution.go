package solution

// Runtime 4 ms
// Memory 2.5 MB

// TODO: overflow handling is hacky, need to fix it

func reverse(x int) int {
	negative := false
	if x == 0 {
		return 0
	}
	if x < 0 {
		x = -x
		negative = true
	}

	output := 0

	for x > 0 {
		digit := x % 10
		output = output*10 + digit
		x = x / 10
	}

	if negative {
		if output > 2147483647 {
			return 0
		}
		return -1 * output
	}
	if output > 2147483648 {
		return 0
	}
	return output

}
