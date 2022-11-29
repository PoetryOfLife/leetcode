package sort

// BubbleSort 冒牌排序，重复扫描需要排序的序列，比较相邻的两个元素，如果他们顺序错误就把他们交换位置，最大的时间复杂度为O(n^2)，空间复杂度为O(1)
func BubbleSort(ary []int) []int {
	for i := range ary {
		for j := 0; j < len(ary)-1-i; j++ {
			if ary[j] > ary[j+1] {
				ary[j], ary[j+1] = ary[j+1], ary[j]
			}
		}
	}
	return ary
}

// SelectSort 选择排序，冒泡升级版，旨在把数组中最大（最小）一一选择出来放在前面，时间复杂度为O(n^2)，空间复杂度为O(1)
func SelectSort(ary []int) []int {
	length := len(ary)
	for i := 0; i < length; i++ {
		index := i
		for j := i; j < length; j++ {
			if ary[j] < ary[index] {
				index = j
			}
		}
		if index != i {
			ary[index], ary[i] = ary[i], ary[index]
		}
	}
	return ary
}

// InsertionSort 插入排序，构建有序序列，对于未排序序列，在已排序序列中从后往前扫描，找到相应的位置并插入，时间复杂度为O(n^2)，空间复杂度为O(1)
func InsertionSort(ary []int) []int {
	for i := 1; i < len(ary); i++ {
		temp := ary[i]
		index := i
		for index > 0 && temp < ary[index-1] {
			ary[index], ary[index-1] = ary[index-1], ary[index]
			index--
		}
		ary[index] = temp
	}
	return ary
}

// MergeSort 归并排序， 把序列拆成子序列，先把子序列排序，再合并，通常使用递归的方式，时间复杂度O(nlogn)，空间复杂度O(n)
func MergeSort(ary []int) []int {
	temp := make([]int, len(ary))
	internalMerge(ary, temp, 0, len(ary)-1)
	return ary
}

func internalMerge(ary, temp []int, left, right int) {
	if left < right {
		middle := (left + right) / 2
		internalMerge(ary, temp, left, middle)
		internalMerge(ary, temp, middle+1, right)
		mergeAry(ary, temp, left, middle, right)
	}
}

func mergeAry(ary, temp []int, left, middle, right int) {
	i, j, k := left, right, 0
	for i <= middle && j <= right {
		if ary[i] <= ary[j] {
			temp[k] = ary[i]
			k++
			i++
		} else {
			ary[k] = ary[j]
			k++
			j++
		}
	}
	for i <= middle {
		temp[k] = ary[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = ary[j]
		j++
		k++
	}
	for i2 := 0; i2 < k; i2++ {
		ary[left+i2] = temp[i2]
	}
}

// QuickSort 快速排序，选定一个基准值，然后根据这个基准值分为两部分，然后一次对这两个部分进行递归处理，时间复杂度O(nlogn)，空间复杂度O(nlogn)
func QuickSort(ary []int) []int {
	qSort(ary, 0, len(ary)-1)
	return ary
}

func qSort(ary []int, low, high int) {
	if low > high {
		return
	}
	pivot := partition(ary, low, high)
	qSort(ary, low, pivot-1)
	qSort(ary, pivot+1, high)
}

func partition(ary []int, low, high int) int {
	pivot := ary[low]
	for low < high {
		for low < high && ary[high] >= pivot {
			high--
		}
		ary[low] = ary[high]
		for low < high && ary[low] <= pivot {
			low++
		}
		ary[high] = ary[low]
	}
	ary[low] = pivot
	return low
}

// ShellSort 希尔排序
func ShellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}
