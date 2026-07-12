package main

func main() {

}

// 这是力扣hot100的第一题，当时熟悉go语法的时候写了一下，时间复杂度为O(n)
func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for i, num := range nums {
		maps[num] = i
	}
	//不能在map中遍历
	for i, num := range nums {
		need := target - num
		if idx, ok := maps[need]; ok && idx != i {
			return []int{i, idx}
		}
	}
	return nil
}
