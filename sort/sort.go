package sort

// BubbleSort 冒牌排序，核心在于相邻两个数字比较交换位置，时间复杂度为O(n^2)
func BubbleSort(ary []int) []int {
	for i := len(ary) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if ary[j] > ary[j+1] {
				ary[j+1], ary[j] = ary[j], ary[j+1]
			}
		}
	}
	return ary
}

// SelectSort 选择排序，冒泡升级版，旨在把数组中最大（最小）一一选择出来放在前面
func SelectSort(ary []int) []int {
	var min int
	for i := 0; i < len(ary)-1; i++ {
		min = i
		for j := i + 1; j < len(ary); j++ {
			if ary[min] > ary[j] {
				min = j
			}
		}
		if min != i {
			ary[min], ary[i] = ary[i], ary[min]
		}
	}
	return ary
}

// InsertionSort 插入排序，构建有序序列，对于未排序序列，在已排序序列中从后往前扫描，找到相应的位置并插入
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

// MergeSort 把序列拆成子序列，先把子序列排序，再合并，通常使用递归的方式
func MergeSort(ary []int) []int {
	temp := make([]int, len(ary))
	internalMergeSort(ary, temp, 0, len(ary)-1)
	return ary
}

func internalMergeSort(ary, temp []int, left, right int) {
	if left < right {
		middle := (left + right) / 2
		internalMergeSort(ary, temp, left, middle)
		internalMergeSort(ary, temp, middle+1, right)
		mergeSortedArray(ary, temp, left, middle, right)
	}
}
func mergeSortedArray(ary, temp []int, left, middle, right int) {
	i, j, k := left, middle+1, 0

	for i <= middle && j <= right {
		if ary[i] < ary[j] {
			temp[k] = ary[i]
			i++
		} else {
			temp[k] = ary[j]
			j++
		}
		k++
	}
	for i <= middle {
		temp[k] = ary[i]
		k++
		i++
	}
	for j <= right {
		temp[k] = ary[j]
		k++
		j++
	}
	for index := 0; index < k; index++ {
		ary[left+index] = temp[index]
	}
}
