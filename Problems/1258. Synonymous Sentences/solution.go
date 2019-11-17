package solution

import "strings"
import "sort"

// Runtime 0 ms, faster than ?%
// Memory 4 MB, less than ?%

func generateSentences(synonyms [][]string, text string) []string {
	m := make(map[string][]string)
	for _, s := range synonyms {
		m[s[0]] = append(m[s[0]], s[1])
		m[s[1]] = append(m[s[1]], s[0])
	}

	queue := []string{text}
	result := make(map[string]struct{})

	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]

		result[t] = struct{}{}

		words := strings.Split(t, " ")
		for i := 0; i < len(words); i++ {
			if _, ok := m[words[i]]; !ok {
				continue
			}

			for _, s := range m[words[i]] {
				newWords := append(append(words[:i], s), words[i+1:]...)
				newText := strings.Join(newWords, " ")
				if _, ok := result[newText]; !ok {
					queue = append(queue, newText)
				}
			}
		}
	}

	res := make([]string, 0, len(result))
	for k := range result {
		res = append(res, k)
	}
	sort.Sort(sort.StringSlice(res))
	return res

}
