package solution

// Runtime 12 ms
// Memory 6.2 MB

func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make([]course, numCourses)
	for _, p := range prerequisites {
		courses[p[0]].pre = append(courses[p[0]].pre, p[1])
	}

	// fmt.Println(courses)

	for i, c := range courses {
		if c.visited {
			continue
		}

		if isCycle(i, courses) {
			return false
		}
	}
	// fmt.Println(courses)

	return true
}

func isCycle(i int, courses []course) bool {
	courses[i].visiting = true
	for _, p := range courses[i].pre {
		if courses[p].visiting {
			return true
		}
		courses[p].visiting = true
		if isCycle(p, courses) {
			return true
		}
		courses[p].visiting = false
	}
	courses[i].visiting = false
	courses[i].visited = true
	return false
}

type course struct {
	pre      []int
	visiting bool
	visited  bool
}
