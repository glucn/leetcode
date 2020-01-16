package solution

// Runtime 0 ms, faster than 100%
// Memory 2.9 MB, less than 100%
// TODO: it is given that the array is sorted, make use of that condition

func nextGreatestLetter(letters []byte, target byte) byte {
    result := byte('z' + 1)
    smallest := byte('z' + 1)
    for _, c := range letters {
        if c > target && c < result {
            result = c
        }
        if c < smallest {
            smallest = c
        }
    }
    if result == byte('z' + 1) {
        return smallest
    }
    return result
}