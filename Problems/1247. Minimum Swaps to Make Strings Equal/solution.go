package solution

// Runtime 0 ms
// Memory 2.1 MB

func minimumSwap(s1 string, s2 string) int {
	mx := make(map[int]struct{})
	my := make(map[int]struct{})

	for i := range s1 {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				mx[i] = struct{}{}
			} else {
				my[i] = struct{}{}
			}
		}
	}
	// fmt.Println(len(mx), len(my))

	if (len(mx)+len(my))%2 == 1 {
		return -1
	}

	if len(mx)%2 == 0 && len(my)%2 == 0 {
		return len(mx)>>1 + len(my)>>1
	}
	return len(mx)>>1 + len(my)>>1 + 2
}
