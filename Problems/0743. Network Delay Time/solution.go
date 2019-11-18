package solution

import "math"
import "container/heap"

// Runtime 68 ms, faster than 53.06%
// Memory 6.9 MB, less than 100%

// Dijkstra's Algorithm
func networkDelayTime(times [][]int, N int, K int) int {
	// TODO: validate K < N
	graph := make([][]edge, N+1)
	for i := range graph {
		graph[i] = make([]edge, 0, N+1)
	}

	// add edges to graph
	for _, t := range times {
		if len(t) < 3 {
			return -1
		}
		// TODO: validate t[0], t[1] < N
		e := edge{to: t[1], time: t[2]}
		graph[t[0]] = append(graph[t[0]], e)
	}

	time := make([]int, N+1)
	visited := make([]bool, N+1)
	for i := range time {
		time[i] = math.MaxInt32
	}
	time[K] = 0

	h := &pathHeap{
		paths: &pathSlice{},
	}
	heap.Push(h.paths, path{index: K, time: 0})

	for len(*h.paths) > 0 {
		current := heap.Pop(h.paths).(path)
		if visited[current.index] {
			continue
		}

		for _, e := range graph[current.index] {
			if !visited[e.to] {
				newPath := path{index: e.to, time: current.time + e.time}
				heap.Push(h.paths, newPath)
			}
		}

		time[current.index] = current.time
		visited[current.index] = true
	}

	max := 0
	for i := 1; i <= N; i++ {
		if time[i] == math.MaxInt32 {
			return -1
		}
		if time[i] > max {
			max = time[i]
		}
	}
	return max
}

type edge struct {
	to   int
	time int
}

type path struct {
	index int
	time  int
}

type pathHeap struct {
	paths *pathSlice
}

type pathSlice []path

func (ps pathSlice) Len() int           { return len(ps) }
func (ps pathSlice) Less(i, j int) bool { return ps[i].time < ps[j].time }
func (ps pathSlice) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }
func (ps *pathSlice) Push(x interface{}) {
	*ps = append(*ps, x.(path))
}

func (ps *pathSlice) Pop() interface{} {
	old := *ps
	p := old[len(old)-1]
	*ps = old[:len(old)-1]
	return p
}
