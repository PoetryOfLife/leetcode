package find

func Bsearch(arr []int, n, value int) int {
	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BsearchRecursion(arr []int, low, high, value int) int {
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1) // + 的优先级高于 >>
	if arr[mid] == value {
		return mid
	} else if arr[mid] < value {
		return BsearchRecursion(arr, mid+1, high, value)
	} else {
		return BsearchRecursion(arr, low, mid-1, value)
	}
}

func Bsearch2(arr []int, n, value int) int {
	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if (mid == 0) || (arr[mid-1] != value) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func Bsearch3(arr []int, n, value int) int {
	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if (mid == n-1) || (arr[mid+1] != value) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func Bsearch4(arr []int, n, value int) int {
	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] >= value {
			if (mid == 0) || (arr[mid-1] < value) {
				return mid
			} else {
				high = mid - 1
			}
		} else if arr[mid] < value {
			low = mid + 1
		}
	}
	return -1
}

func Bsearch5(arr []int, n, value int) int {
	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > value {
			high = mid - 1

		} else if arr[mid] <= value {
			if (mid == n-1) || (arr[mid+1] > value) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
