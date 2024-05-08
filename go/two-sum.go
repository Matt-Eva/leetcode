func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    var solution []int

    for idx, val := range nums {
        i, ok := m[target - val]
        if ok {
            solution = []int{idx, i}
        } 
        m[val] = idx
    }

    return solution
}