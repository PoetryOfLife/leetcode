package question

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int, 0)
	for k, v := range nums {
		if x, ok := hashTable[target-v]; ok {
			return []int{x, k}
		}
		hashTable[v] = k
	}
	return nil
}
