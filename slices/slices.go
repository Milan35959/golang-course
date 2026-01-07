package main

// slices are like dynamic arrays
// most used construct in go
// +useful methods
func main() {
	//uninitialized slice is nil
	// var nums []int

	// // fmt.Println(nums == nil) //tru

	// fmt.Println(len(nums))

	// var nums = make([]int, 2, 5)
	// capacity-> maximum numbers of elements  can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)

	// nums := []int{}

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// nums = append(nums, 9)
	// nums = append(nums, 10)

	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// //copy function

	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//slice operators

	// var nums = []int{1, 2, 3, 4, 5}

	// fmt.Println(nums[0:1]) //start index inclusive, end index exclusive
	// fmt.Println(nums[:1]) //from start to index 1(exclusive)
	// fmt.Println(nums[1:])

	//slice package

	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2))

	//2d slices

	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)
}
