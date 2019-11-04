package solution

// Runtime 64 ms, faster than 5.98%
// Memory 5.7 MB, less than 87.5%
// TODO: improve the solution

func firstUniqChar_v3(s string) int {
	m := make(map[byte]data)

	for i := range s {
		if d, ok := m[s[i]]; ok {
			m[s[i]] = data{count: d.count + 1, first: d.first}
		} else {
			m[s[i]] = data{count: 1, first: i}
		}
	}

	min := len(s)
	for _, v := range m {
		if v.count == 1 && v.first < min {
			min = v.first
		}
	}
	if min == len(s) {
		return -1
	}
	return min
}

type data struct {
	count int
	first int
}
