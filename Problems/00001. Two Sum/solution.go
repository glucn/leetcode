package solution

// Runtime 4ms
// Memory 3.5 MB

func twoSum(nums []int, target int) []int {
    complimentryMap := make(map[int]int, 0)
    for i, n := range nums {
        if c, ok := complimentryMap[target-n]; ok {
            return []int{c, i}
        }
        complimentryMap[n] = i
    }
    
    return nil
}