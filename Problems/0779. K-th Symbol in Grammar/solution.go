package solution

// Runtime 0 ms, faster than 100%
// Memory 2 MB, less than 100%

func kthGrammar(N int, K int) int {
	return kthGrammarZeroIndexed(N-1, K-1)
}

func kthGrammarZeroIndexed(N int, K int) int {
	if N == 0 {
		return 0
	}
	d := kthGrammarZeroIndexed(N-1, K>>1)
	if d == 0 {
		if K%2 == 0 {
			return 0
		}
		return 1
	}
	if K%2 == 0 {
		return 1
	}
	return 0
}
