## goal is to get all consecutive combinations within a string.

def consecutive_word_in_string(string: str) -> dict[str, str]: 
    print("string length", len(string))
    word_map : dict[str, str] = {}
    for i in range(len(string)):
        char = string[i]
        n = i
        while n < len(string):
            slc = string[i:n+1]
            word_map[slc] = slc
            n += 1
    print("entry number", len(word_map))
    # print(word_map)
    return word_map

consecutive_word_in_string("hello world it's me margaret")