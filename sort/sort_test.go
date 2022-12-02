package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println(BubbleSort([]int{-23, -34, -5, -23, 23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, -234, -234, 234}))
}

func TestSelectSort(t *testing.T) {
	fmt.Println(SelectSort([]int{-23, -34, -5, -23, 23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, -234, -234, 234}))
}

func TestInsertionSort(t *testing.T) {
	fmt.Println(InsertionSort([]int{-23, -34, -5, -23, 23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, -234, -234, 234}))
}

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort([]int{-23, -34, -5, -23, 23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, -234, -234, 234}))
}

func TestShellSort(t *testing.T) {
	fmt.Println(ShellSort([]int{-23, -34, -5, -23, 23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, -234, -234, 234}))
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{-23, -34, -5, -23, 23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, -234, -234, 234}))
}

func TestHeapSort(t *testing.T) {
	fmt.Println(HeapSort([]int{-23, -34, -5, -23, 23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, -234, -234, 234}))
}

func TestCountingSort(t *testing.T) {
	fmt.Println(CountingSort([]int{23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, 234}, 20000))
}

func TestBucketSort(t *testing.T) {
	fmt.Println(BucketSort([]int{23, 56, 1234, 63, 56, 734, 36, 24, 2334, 2, 24, 624, 12343, 2482, 234}, 5))
}
