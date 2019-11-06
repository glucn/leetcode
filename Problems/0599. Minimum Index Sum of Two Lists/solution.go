package solution

// Runtime 36 ms, faster than 57.97% // TODO: improve this
// Memory 7.1 MB, less than 100%

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]poi)
	for i, name := range list1 {
		m[name] = poi{idx1: i, idx2: -1}
	}
	for i, name := range list2 {
		if mm, ok := m[name]; ok {
			m[name] = poi{idx1: mm.idx1, idx2: i}
		}
	}

	min := len(list1) + len(list2)
	var minNames []string
	for k, v := range m {
		if v.idx2 != -1 {
			i := v.idx1 + v.idx2
			if i < min {
				min = i
				minNames = []string{k}
			} else if i == min {
				minNames = append(minNames, k)
			}
		}
	}

	return minNames
}

type poi struct {
	idx1, idx2 int
}
