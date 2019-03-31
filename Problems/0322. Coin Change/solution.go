package solution

// Runtime 184 ms
// Memory 9.9 MB

import "math"
import "sort"

func coinChange(coins []int, amount int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	calculated := make(map[int]int, amount)

	return subFunc(coins, amount, calculated)
}

func subFunc(coins []int, amount int, calculated map[int]int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if s, ok := calculated[amount]; ok {
		return s
	}

	best := math.MaxInt32
	for _, c := range coins {
		s := subFunc(coins, amount-c, calculated)
		if s != -1 && s < best {
			best = s + 1
		}

	}
	if best == math.MaxInt32 {
		calculated[amount] = -1
		return -1
	}
	calculated[amount] = best
	return best
}
