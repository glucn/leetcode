package solution

// Runtime 16 ms, faster than 18.64%
// Memory 5.3 MB, less than 100%
// TODO: leverage the fact that the input is sorted to make it faster

func findSpecialInteger(arr []int) int {    
    target := len(arr) / 4
    current := arr[0]
    count := 0
    
    for _, n := range arr {
        if n == current {
            count++
            if count > target {
                return n
            }
        } else {
            current = n
            count = 1
        }
    }
    return -1
}
