func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    result := make([][]string, 0, len(m))

    for _ , v := range strs {
        r := []rune(v)
        sort.Slice(r, func(i, j int) bool {
            return r[i] < r[j]
        })
        s := string(r)

        val, ok := m[s]
        if ok {    
            m[s] = append(val, v)
        } else{
            m[s] = []string{v}
        }
    }

    for _, v := range m{
       result = append(result, v)
    }

    return result
}