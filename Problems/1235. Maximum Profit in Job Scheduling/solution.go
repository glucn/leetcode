package solution

import "sort"

// Runtime 176 ms
// Memory 9.5 MB

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	if len(startTime) == 1 {
		return profit[0]
	}

	var jobs []job
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, job{start: startTime[i], end: endTime[i], profit: profit[i]})
	}

	sort.Sort(jobArray(jobs))

	dpProfit := []int{0}
	dpEndTime := []int{1}

	for _, j := range jobs {
		// fmt.Println(j)

		index := smallerOrEqual(dpEndTime, j.start)
		// fmt.Println("output", index)

		total := dpProfit[index] + j.profit
		if total > dpProfit[len(dpProfit)-1] {
			dpEndTime = append(dpEndTime, j.end)
			dpProfit = append(dpProfit, total)
		}

		// fmt.Println(dpEndTime)
		// fmt.Println(dpProfit)
	}

	return dpProfit[len(dpProfit)-1]
}

func smallerOrEqual(times []int, t int) int {
	i := sort.Search(len(times), func(n int) bool { return times[n] > t })
	return max(i-1, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type job struct {
	start, end, profit int
}

type jobArray []job

func (h jobArray) Len() int           { return len(h) }
func (h jobArray) Less(i, j int) bool { return h[i].end < h[j].end }
func (h jobArray) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
