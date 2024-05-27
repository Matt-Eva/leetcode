from encode_decode import Solution

def test_encode_decode():
    list_case = ["abc", "dfg", "hijk", "lmnopqrstuv"]
    string_case = "abcdfghijklmnopqrstuv"
    solution = Solution()
    encoded = solution.encode(["abc", "dfg", "hijk", "lmnopqrstuv"])
    assert encoded == string_case
    decoded = solution.decode(encoded)
    for i in range(len(decoded)):
        assert decoded[i] == list_case[i] 