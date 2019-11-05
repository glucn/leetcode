package solution

// Runtime 0 ms, faster than 100%
// Memory 2 MB, less than 100%

func isHappy(n int) bool {
	slow, fast := n, next(n)
	for fast != 1 && slow != fast { // we don't need to worry about fast goes up to infinity
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1
}

func next(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n = n / 10
	}
	return sum
}
