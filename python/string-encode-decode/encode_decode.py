class Solution:

    def encode(self, strs: list[str]) -> str:
        self.encode_list = []
        index_tracker = -1
        for string in strs:
            string_length = len(string)
            index_tracker += string_length
            self.encode_list.append(index_tracker)
        return "".join(strs)

    def decode(self, s: str) -> list[str]:
        starting_index = 0
        decoded = []
        for index in self.encode_list:
            string_slice = s[starting_index : index + 1]
            decoded.append(string_slice)
            starting_index = index + 1
        return decoded