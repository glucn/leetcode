package solution

// Runtime 3704 ms
// Memory 9.2 MB

func circularPermutation(n int, start int) []int {
	temp := grayCode(n)
	for i := range temp {
		if temp[i] == start {
			return append(temp[i:], temp[:i]...)
		}
	}
	// should not get to here
	return nil
}

func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}
	firstHalf := grayCode(n - 1)
	secondHalf := make([]int, len(firstHalf))
	leading := 1 << uint(n-1)
	for i := range firstHalf {
		secondHalf[i] = leading + firstHalf[len(firstHalf)-i-1]
	}
	return append(firstHalf, secondHalf...)
}
