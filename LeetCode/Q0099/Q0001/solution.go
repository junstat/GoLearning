package Q0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		b := target - a
		if j, exist := m[b]; exist {
			return []int{j, i}
		}
		m[a] = i
	}
	return []int{}
}
