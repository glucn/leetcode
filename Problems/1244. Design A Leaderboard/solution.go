package solution

import "container/heap"

// Runtime 28 ms
// Memory 6.3 MB

type Leaderboard struct {
	pq PriorityQueue
}

func Constructor() Leaderboard {
	return Leaderboard{
		pq: make([]*player, 0),
	}
}

func (l *Leaderboard) AddScore(playerId int, score int) {
	for _, p := range l.pq {
		if p.ID == playerId {
			l.pq.update(p, p.score+score)
			return
		}
	}
	heap.Push(&l.pq, &player{ID: playerId, score: score})
}

func (l *Leaderboard) Top(K int) int {
	sum := 0
	var waiting []*player
	for i := 0; i < K; i++ {
		p := heap.Pop(&l.pq)
		// fmt.Println("Top", p)
		if p == nil {
			break
		}
		sum += p.(*player).score
		waiting = append(waiting, p.(*player))
	}
	for _, w := range waiting {
		heap.Push(&l.pq, &player{ID: w.ID, score: w.score})
	}
	return sum
}

func (l *Leaderboard) Reset(playerId int) {
	for _, p := range l.pq {
		if p.ID == playerId {
			l.pq.update(p, 0)
			return
		}
	}
}

type player struct {
	ID    int
	score int
	index int
}

type PriorityQueue []*player

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].score > pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	// fmt.Println(x)
	n := len(*pq)
	item := x.(*player)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *player, score int) {
	item.score = score
	heap.Fix(pq, item.index)
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
