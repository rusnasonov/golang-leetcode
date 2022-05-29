package algorithms


func twoSum(nums []int, target int) []int {
    t := make(map[int]int)
    var res []int

    for idx, num := range nums {
        exp := target - num
        expIdx, exists := t[exp]
        if exists && expIdx != idx {
            res = append(res, idx, expIdx)
            break 
        }
        t[num] = idx
    }
    return res
}
