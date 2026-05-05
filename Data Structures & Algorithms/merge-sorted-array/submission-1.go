func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p3 := m-1, n-1, len(nums1) - 1
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
			p3--
			continue
		}

		nums1[p3] = nums2[p2]
		p2--
		p3--
	}
}
