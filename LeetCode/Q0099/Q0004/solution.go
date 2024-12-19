package Q0004

func find(n1 []int, i int, n2 []int, j int, k int) int {
	if len(n1)-i > len(n2)-j {
		return find(n2, j, n1, i, k)
	}
	if i >= len(n1) {
		return n2[j+k-1]
	}
	if k == 1 {
		return min(n1[i], n2[j])
	} else {
		si := min(i+k/2, len(n1))
		sj := j + k - k/2
		if n1[si-1] > n2[sj-1] {
			return find(n1, i, n2, sj, k-(sj-j))
		} else {
			return find(n1, si, n2, j, k-(si-i))
		}
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b

	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tot := len(nums1) + len(nums2)
	if tot%2 == 0 {
		left := find(nums1, 0, nums2, 0, tot/2)
		right := find(nums1, 0, nums2, 0, tot/2+1)
		return float64(left+right) / 2
	}
	return float64(find(nums1, 0, nums2, 0, tot/2+1))
}
