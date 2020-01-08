package solution

// Runtime 0 ms, faster than 100%
// Memory 1.9 MB, less than 100%

func subtractProductAndSum(n int) int {
	p, s := getProductAndSum(n)
    return p-s
}

func getProductAndSum(n int) (int, int) {
    if n < 10 {
        return n, n
    }
    p, s := getProductAndSum(n/10)
    d := n %10
    return p*d, s+d
}