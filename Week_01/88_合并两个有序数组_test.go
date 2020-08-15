package Week_01

import "sort"

// 先合并再排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

// 双指针，从后向前插入；最后的数最大
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 { //当nums1和nums2都有数据时，从最大下标处开始比较。大的放到nums1的最右方
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m-- // nums1大，则m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n-- // nums2大，则n--
		}
	}
	if m == 0 { //分两种情况：1.nums1原本的有效数据为0；2.经过for循环的比较后，nums1的数据因为大于nums2的数据，所以优先塞到偏右面去了，会导致nums2还有剩余数据
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
}
