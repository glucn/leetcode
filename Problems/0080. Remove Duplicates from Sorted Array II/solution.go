package solution

// Runtime 8 ms
// Memory 6 MB

func removeDuplicates(nums []int) int {
    if len(nums) < 1 {
        return 0
    }
    
    count := 0
    
    for i := 0 ; i < len(nums); i++ {
        if i == len(nums) - 1 || nums[i] != nums[i+1] {
            // fmt.Println(count, i, nums[i])
            nums[count] = nums[i]
            count++
            continue
        }
        if nums[i] == nums[i+1] && (i == len(nums) - 2 || nums[i] != nums[i+2]) {

            nums[count] = nums[i]
            nums[count+1] = nums[i]
            count += 2
            i++
            continue
        }
        
        j := i+1
        for j< len(nums) && nums[i] == nums[j] {
            j++
        }

        nums[count] = nums[i]
        nums[count+1] = nums[i]
        count += 2    
        i = j - 1
    }
    
    return count
}