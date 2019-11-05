package solution

// Runtime 0 ms, faster than 100%
// Memory 2 MB, less than 100%


func circularArrayLoop(nums []int) bool {
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		slow, fast := i, next(i, nums)

		// same direction, and none of both is 0
		for nums[fast]*nums[i] > 0 && nums[next(fast, nums)]*nums[i] > 0 {
			fast = next(next(fast, nums), nums)
			slow = next(slow, nums)

			if fast == slow {
				if slow == next(slow, nums) {
					// loop with size 1
					break
				}
				return true
			}
		}

		// no loop found from i
		j := i
		n := nums[j]
		for nums[j]*n > 0 {
			nums[j], j = 0, next(j, nums)
		}
	}
	return false
}

func next(i int, nums []int) int {
	n := (i + nums[i]) % len(nums)
	if n >= 0 {
		return n
	}
	return n + len(nums)
}
