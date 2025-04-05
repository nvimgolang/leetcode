package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	pointer := m + n - 1 // Указатель на место записи в nums1
	array1 := m - 1      // Указатель на последний элемент в nums1
	array2 := n - 1      // Указатель на последний элемент в nums2

	for array2 >= 0 {
		if array1 >= 0 && nums1[array1] > nums2[array2] {
			nums1[pointer] = nums1[array1]
			array1--
		} else {
			nums1[pointer] = nums2[array2]
			array2--
		}
		pointer--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{4, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1) // [1 2 3 4 5 6]
}
