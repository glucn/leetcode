package solution

// Runtime 0 ms, faster than 100%
// Memory 2.3 MB, less than 100%

func sortColors(nums []int) {
	idx0, idx1, idx2 := -1, -1, -1
	for i := range nums {
		switch nums[i] {
		case 0:
			nums[idx0+1] = 0
			if idx1 != idx0 {
				nums[idx1+1] = 1
			}
			if idx2 != idx1 {
				nums[idx2+1] = 2
			}
			idx0++
			idx1++
			idx2++

		case 1:
			nums[idx1+1] = 1
			if idx2 != idx1 {
				nums[idx2+1] = 2
			}
			idx1++
			idx2++

		case 2: // maybe can replace idx2 with i
			nums[idx2+1] = 2
			idx2++
		}
	}

}
