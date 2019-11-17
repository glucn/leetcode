package solution

// Runtime 52 ms, faster than ?%
// Memory 7.1 MB, less than ?%

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	m := make(map[string]*node)
	noroot := make(map[string]struct{})

	for _, r := range regions {
		if mm, ok := m[r[0]]; ok {
			mm.put(r[1:], m, noroot)
		} else {
			n := &node{region: r[0]}
			noroot[r[0]] = struct{}{}
			m[r[0]] = n
			n.put(r[1:], m, noroot)
		}
	}

	if len(noroot) != 1 {
		return ""
	}

	var root string
	for k := range noroot {
		root = k
	}

	_, _, p := dfs(m[root], region1, region2)
	return p
}

type node struct {
	region string
	next   []*node
}

func (n *node) put(children []string, regionMap map[string]*node, noRootMap map[string]struct{}) {
	for _, r := range children {
		if _, ok := noRootMap[r]; ok {
			delete(noRootMap, r)
		}
		if nn, ok := regionMap[r]; ok {
			n.next = append(n.next, nn)
		} else {
			newNode := &node{region: r}
			n.next = append(n.next, newNode)
			regionMap[r] = newNode
		}
	}
}

func dfs(node *node, r1, r2 string) (bool, bool, string) {
	if node == nil {
		return false, false, ""
	}

	found1 := node.region == r1
	found2 := node.region == r2
	for _, nn := range node.next {
		f1, f2, parent := dfs(nn, r1, r2)
		if parent != "" {
			return true, true, parent
		}
		if f1 {
			found1 = f1
		}
		if f2 {
			found2 = f2
		}
	}
	if found1 && found2 {
		return true, true, node.region
	}
	return found1, found2, ""
}
