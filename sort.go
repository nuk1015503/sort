package sort

// BubbleSort bubble sort
func BubbleSort(arr []int) []int {
	if arr == nil {
		return arr
	}

	last := len(arr) - 1

	for range arr {
		for idx := 0; idx < last; idx++ {
			if arr[idx] > arr[idx+1] {
				tmp := arr[idx]
				arr[idx] = arr[idx+1]
				arr[idx+1] = tmp
			}
		}
		last--
	}
	return arr
}

// SelectionSort selection sort
func SelectionSort(arr []int) []int {
	if arr == nil {
		return arr
	}

	first := 0

	for range arr {
		min := arr[first]
		tmpIndex := first
		for idx := first; idx < len(arr); idx++ {
			if arr[idx] < min {
				min = arr[idx]
				tmpIndex = idx
			}
		}
		tmp := arr[tmpIndex]
		arr[tmpIndex] = arr[first]
		arr[first] = tmp

		first++
	}
	return arr
}

// InsertionSort insertion sort
func InsertionSort(arr []int) []int {
	if arr == nil {
		return arr
	}

	first := 0

	for range arr {
		for idx := first; idx > 0; idx-- {
			if arr[idx] < arr[idx-1] {
				tmp := arr[idx]
				arr[idx] = arr[idx-1]
				arr[idx-1] = tmp
			}
		}

		first++
	}
	return arr
}
