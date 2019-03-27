package solution

// Runtime 4 ms
// Memory 4.7 MB

import "strings"

func licenseKeyFormatting(S string, K int) string {
    outputSlice := make([]byte, 2*len(S))
    var counter int
    var length int
    for i := len(S)-1 ; i >=0 ; i-- {
        if S[i] == '-' {
            continue
        }
        if counter % K == 0 && length != 0 {
            outputSlice[length] = '-'
            length++
        }
        outputSlice[length] = S[i]
        length++
        counter++
        
    } 
    
    for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
        outputSlice[i], outputSlice[j] = outputSlice[j], outputSlice[i]
    }
    
    return strings.ToUpper(string(outputSlice[0:length]))
}