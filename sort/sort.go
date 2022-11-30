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

// ShellSort 希尔排序，把记录按下表的一定增量分组，对每组使用直接插入排序算法排序；随着增量逐渐减少，
//每组包含的关键词越来越多，当增量减至1时，整个文件恰被分成一组，算法便终止，时间复杂度O(nlogn)，空间复杂度O(1)
func ShellSort(ary []int) []int {
	length := len(ary)
	temp, gap := 0, length/2
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp = ary[i]
			j := i - gap
			for j >= 0 && ary[j] > temp {
				ary[j+gap] = ary[j]
				j = j - gap
			}
			ary[j+gap] = temp
		}
		gap /= 2
	}
	return ary
}

// MergeSort 归并排序，把序列拆成子序列，先把子序列排序，再合并，通常使用递归的方式，时间复杂度O(nlogn)，空间复杂度O(n)
func MergeSort(ary []int) []int {
	length := len(ary)
	if length < 2 {
		return ary
	}
	middle := length / 2
	left := ary[0:middle]
	right := ary[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return
}

// QuickSort 快速排序，选定一个基准值，把一个序列分成两部分，其中一部分均比另一部分小，然后通过递归的方式分别对两个部分进行相同方式的排序，直到整个序列有序，时间复杂度O(nlogn)，空间复杂度O(nlogn)
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
