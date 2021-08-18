package array

func mergeArr(nums1 []int, m int, nums2 []int, n int)  {
	if len(nums1) == 0 || len(nums2) == 0 {
		return
	}

	n1 := m-1
	n2 := n-1
	l := m+n-1

	for n2 >= 0 {
		if n1 >= 0 && nums1[n1] > nums2[n2] {
			nums1[l] = nums1[n1]
			n1 --
		}else {
			nums1[l] = nums2[n2]
			n2 --
		}
		l --
	}
}