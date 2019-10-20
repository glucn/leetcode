package solution

import "strings"
import "sort"

// Runtime 136 ms
// Memory 136.9 MB

func removeSubfolders(folder []string) []string {
	sort.Sort(sort.StringSlice(folder))
	var result []string

	for _, f := range folder {
		if len(result) == 0 {
			result = append(result, f)
			continue
		}
		if shouldAdd(f, result[len(result)-1]) {
			result = append(result, f)
			continue
		}
	}

	return result

}

func shouldAdd(this, last string) bool {
	if len(this) < len(last) {
		return true
	}
	if strings.HasPrefix(this, last) {
		if len(last) == len(this) {
			return false
		}
		if this[len(last)] == '/' {
			return false
		}
		return true
	}
	return true
}
