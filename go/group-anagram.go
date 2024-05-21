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

	// Optimized solution - not my own:

	// func hash(s string) string {
	// 	res := make([]byte, 26)
	// 	for _, c := range s {
	// 		res[c-'a'] += 1
	// 	}
	// 	return string(res)
	// }
	
	// func groupAnagrams(strs []string) [][]string {
	// 	res := [][]string{}
	// 	m := make(map[string]int)
	// 	for _, w := range strs {
	// 		h := hash(w)
	// 		idx, ok := m[h]
	// 		if ok {
	// 			res[idx] = append(res[idx], w)
	// 		} else {
	// 			res = append(res, []string{w})
	// 			m[h] = len(res) - 1
	// 		}
	// 	}
	
	// 	return res
	// }
}