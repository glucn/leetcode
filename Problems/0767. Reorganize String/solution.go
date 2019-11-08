package solution

// Runtime 0 ms, faster than 100%
// Memory 2.2 MB, less than 100%

func reorganizeString(S string) string {
	m := make(map[byte]int)

	for i := range S {
		m[S[i]] = m[S[i]] + 1
	}

	h := NewHeap()

	for k, v := range m {
		h.push(&node{letter: k, count: v})
	}

	result := make([]byte, len(S))
	var tmp *node
	for i := range result {
		n := h.pop()
		if n == nil {
			return ""
		}
		result[i] = n.letter
		if tmp != nil {
			h.push(tmp)
		}
		if n.count > 1 {
			tmp = &node{letter: n.letter, count: n.count - 1}
		} else {
			tmp = nil
		}
	}

	return string(result)

}

type heap struct {
	nodes []*node
}

type node struct {
	letter byte
	count  int
}

func NewHeap() *heap {
	return &heap{
		nodes: nil,
	}
}

func (h *heap) push(n *node) {
	h.nodes = append(h.nodes, n)
	curr := len(h.nodes) - 1
	for parent(curr) >= 0 && h.nodes[curr].count > h.nodes[parent(curr)].count {
		h.nodes[curr], h.nodes[parent(curr)] = h.nodes[parent(curr)], h.nodes[curr]
		curr = parent(curr)
	}
}

func (h *heap) pop() *node {
	if len(h.nodes) == 0 {
		return nil
	}
	n := h.nodes[0]
	h.nodes = h.nodes[1:]
	h.fix(0)
	return n
}

func (h *heap) fix(idx int) {
	if left(idx) < len(h.nodes) && h.nodes[idx].count < h.nodes[left(idx)].count {
		h.nodes[idx], h.nodes[left(idx)] = h.nodes[left(idx)], h.nodes[idx]
		h.fix(left(idx))
	}

	if right(idx) < len(h.nodes) && h.nodes[idx].count < h.nodes[right(idx)].count {
		h.nodes[idx], h.nodes[right(idx)] = h.nodes[right(idx)], h.nodes[idx]
		h.fix(right(idx))
	}

}

func parent(idx int) int {
	return (idx - 1) >> 1
}

func left(idx int) int {
	return 2*idx + 1
}

func right(idx int) int {
	return 2*idx + 2
}
