func removeDuplicates(nums []int) int {
    if (len(nums) == 0 || len(nums) == 1) {
        return len(nums)
    }
    x, lenn := 1, len(nums)
    for i := 1; i < lenn ; i++ {
        if(nums[i] != nums[i - 1]) {
            
            nums[x] = nums[i]
            x++
        }
    }
    return x
}
