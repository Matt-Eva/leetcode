class Solution:

    def recursiveCheck(self, nums, i, numsLen):
        if i + 1 < numsLen and nums[i] != nums[i + 1]:
            print(nums[i])
            nums[i] = nums[i + 1]
            print(nums[i])
            self.recursiveCheck(nums, i + 1, numsLen)
        elif i + 1 < numsLen:
            print("elif")
            nums[i] = nums[i + 1]

    def removeDuplicates(self, nums: List[int]) -> int:
        numsLen = len(nums)
        k = 0
        for i in range(numsLen): 
            k += 1
            n = i + 1
            while n < numsLen and nums[i] == nums[n]:
                if n != i + 1:
                    k -= 1
                    nums[n] = None
                n += 1

            if i > 1 and nums[i] == nums[i - 2]:
                k -=1
        
        for i in range(numsLen):
            if nums[i] == None:
                n = i
                while nums[n] == None:
                    n += 1
                nums[i] = nums[n]
                nums[n]==None
        
        
        # for i in range(numsLen):
        #     print("nums[i]", nums[i])
        #     n = i + 1 
        #     length = 0
        #     while n < numsLen and nums[i] == nums[n]:
        #         length += 1
        #         n += 1

        #     q = i + 1
        #     tracker = length
        #     while q + length < numsLen and tracker > 1:
        #         nums[q + 1] = nums[q + length]
        #         if q + length + 1 < numsLen:
        #             self.recursiveCheck(nums, q + length, numsLen)
        #         q += 1
        #         tracker -= 1

        print(nums)
        return k

# we start with each number by counting the number of duplicates it has
# that number is captured in the "length" variable.
# Once we have the number of duplicates, we know that we need to shift that many numbers to the left in the array
# however, we don't want the array to go out of sorted order, so every number that is shifted left needs to be replaced by the following number
# but, we don't want accidental duplicates, so every time we shift a number over, we need to check if the number following the one we shifted over is the same. If it's not, we need to shift that number over too.
# This requires recursion.

# set duplicates to null. Then, trace forward and set any non-null values to the null values. As soon as non-null values are traced forward, set them to null.