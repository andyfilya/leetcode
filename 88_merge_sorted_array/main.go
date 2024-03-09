package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		nums1Ptr = m - 1
		nums2Ptr = n - 1
		writePtr = m + n - 1
	)

	for writePtr >= 0 && nums2Ptr >= 0 && nums1Ptr >= 0 {
		if nums1[nums1Ptr] > nums2[nums2Ptr] {
			nums1[writePtr] = nums1[nums1Ptr]
			nums1Ptr--
		} else {
			nums1[writePtr] = nums2[nums2Ptr]
			nums2Ptr--
		}
		writePtr--
	}

	for nums1Ptr >= 0 {
		nums1[writePtr] = nums1[nums1Ptr]
		writePtr--
		nums1Ptr--
	}

	for nums2Ptr >= 0 {
		nums1[writePtr] = nums2[nums2Ptr]
		nums2Ptr--
		writePtr--
	}
}
