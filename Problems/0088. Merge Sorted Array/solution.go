package solution

// Runtime 0 ms
// Memory 3.6 MB

func merge(nums1 []int, m int, nums2 []int, n int) {
	for {
		if m == 0 && n == 0 {
			return
		}
		if m == 0 {
			nums1[m+n-1] = nums2[n-1]
			n--
			continue
		}
		if n == 0 {
			nums1[m+n-1] = nums1[m-1]
			m--
			continue
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	return
}
