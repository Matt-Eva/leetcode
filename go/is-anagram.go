package main

func isAnagram(s string, t string) bool {
    // ss := strings.Split(s, "")
    // st := strings.Split(t, "")

    // sort.Strings(ss)
    // sort.Strings(st)

    // return (strings.Join(ss, "") == strings.Join(st, ""))

	//^^ slow solution


	// slen := len(s)
	// if slen != len(t){
	// 	return false
	// }

	// type LetterCount struct {
	// 	T int
	// 	S int
	// }

	// m := make(map[rune]*LetterCount)

	// for i := 0; i < slen; i++  {
	// 	sRune, _ := utf8.DecodeRuneInString(s[i:])
	// 	tRune, _ := utf8.DecodeRuneInString(t[i:])

	// 	sValue := m[sRune]
	// 	if sValue == nil{
	// 		m[sRune] = &LetterCount{S: 1, T: 0}
	// 	} else {
	// 		m[sRune].S += 1
	// 	}

	// 	tValue := m[tRune]
	// 	if tValue == nil {
	// 		m[tRune] = &LetterCount{S: 0, T: 1}
	// 	} else{
	// 		m[tRune].T += 1
	// 	}
	// }


	// for _, v := range m {
	// 	if v.T != v.S{
	// 		return false
	// 	}
	// }

	// return true

	// ^^my solution - better, good with memory, but could be faster
	
	if len(t) != len(s) {
        return false
    }

    letterSet := make(map[rune]int)
    for _, letter := range s {
        letterSet[letter]++
    }

    for _, letter := range t {
       letterSet[letter]--
    }

    for _, count := range letterSet {
        if count != 0 {
            return false
        }
    }

    return true

	// ^^ Best solution according to leetcode - not mine
}