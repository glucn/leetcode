package solution

// Runtime 4 ms
// Memory 7.4 MB

func grayCode(n int) []int {
    length := 1 << uint(n)

    codeMap := make(map[int]struct{}, length)
    output := make([]int, length)
    
    output[0] = 0
    last := 0
    codeMap[0] = struct{}{}
    
    for i := 1; i< length; i++ {
        for j := 0; j < n; j++ {
            x := last ^ (1 << uint(j))
            if _, ok := codeMap[x]; !ok {
                output[i] = x
                codeMap[x] = struct{}{}
                last = x
                break
            }
        }
    }
    
    return output
}