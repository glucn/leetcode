package solution

// Runtime 20 ms, faster than 52.63%
// Memory 7.1 MB, less than 100%

func longestWord(words []string) string {
	root := NewNode()

	for i, w := range words {
		root.put([]byte(w), i)
	}

	queue := []*node{root}
	var result string
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]

		if n.isWord && (len(words[n.index]) > len(result) || (len(words[n.index]) == len(result) && words[n.index] < result)) {
			result = words[n.index]
		}

		for _, nn := range n.next {
			if nn != nil && nn.isWord {
				queue = append(queue, nn)
			}
		}
	}
	return result
}

type node struct {
	index  int
	isWord bool
	next   []*node
}

func NewNode() *node {
	return &node{next: make([]*node, 26)}
}

func (n *node) put(b []byte, index int) {
	if len(b) == 0 {
		n.index = index
		n.isWord = true
		return
	}
	if n.next[b[0]-'a'] == nil {
		n.next[b[0]-'a'] = NewNode()
	}

	n.next[b[0]-'a'].put(b[1:], index)
	return
}
