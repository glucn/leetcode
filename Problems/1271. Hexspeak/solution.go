package solution

import "strconv"

// Runtime 0 ms, faster than ?%
// Memory 2 MB, less than ?%

func toHexspeak(num string) string {
    var result []byte
    numInt, _ := strconv.Atoi(num)
    
    for numInt > 0 {
        digit := numInt % 16
        if digit == 0 {
            result = append(result, 'O')
        } else if digit == 1 {
            result = append(result, 'I')
        } else if digit > 9 {
            result = append(result, ('A' + byte(digit-10)))
        } else {
            return "ERROR"
        }
        numInt = numInt / 16
    }
    return string(revert(result))
}

func revert(b []byte) []byte {
    left, right := 0, len(b)-1
    for left < right {
        b[left], b[right] = b[right], b[left]
        left++
        right--
    }
    return b
}