package array
import "sort"
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	length := len(nums1)
	if length % 2 == 1 {
		return float64(nums1[length/2])
	}

	n1 := len(nums1)/2
	n2 := len(nums1)/2 - 1
	return float64(nums1[n1] + nums1[n2]) / 2
}