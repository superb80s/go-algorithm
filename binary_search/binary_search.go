package binary_search

// BinarySearch 二分查找 升续
func BinarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		midVal := nums[mid]
		if midVal == target {
			return midVal
		} else if midVal > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
