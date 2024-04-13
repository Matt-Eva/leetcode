class Solution:

# new approach - find the start of the bowl and the end of the bowl.
# determine the bowl's final height - either starting or ending height.
# iterate through the base of the bowl, subtracting the floor height from the final height
# add the difference to the total volume
# need to keep track of relative floor maximums and their indexes
    # if we reach the end of an array, but the end is not greater than or equal to the starting height of the bowl, then we need to iterate through the sub array of relative maximums.d
# as we're going through a bowl, we'll need to programmatically remove any relative maximums that are exceeded by a subsequent relative maximum.
    #We can store the relative maximums in a list, along with their specific indexes - so a list of dictionaries.
    #Each time we find a new relative maximum, we append it to the list and iterate through the list to see if preceding relative maximums are less than the current relative maximum.
        # If they are, then we want to remove them from the list.
    #If we are using relative maximums, we will start iterating from the bowl starting index.
# need a few variables
# is_bowl
# starting_height
# starting_index
# relative_maximum keeps track of 


    def trap(self, height):
        i = 0
        total_volume = 0
        is_bowl = False
        bowl_starting_height = 0
        bowl_starting_index = 0
        relative_starting_index = 0
        relative_maximums = []
        for i in range(len(height)):
            # print("index", i)
            # print("total_volume", total_volume)
            # print("is_bowl", is_bowl)
            # print("relative_maximums", relative_maximums)
            if i + 1 >= len(height):
                if is_bowl and height[i] >= bowl_starting_height:
                    for n in range(bowl_starting_index + 1, i):
                        volume = bowl_starting_height - height[n]
                        total_volume += volume

                elif is_bowl and height[i] > height[i - 1]:
                    new_rel_maximums = []
                    for dictionary in relative_maximums:
                        if dictionary.get("height") >= height[i]:
                            new_rel_maximums.append(dictionary)
                    relative_maximums = new_rel_maximums
                    relative_maximums.append({"height": height[i], "index": i})
                    relative_starting_index = bowl_starting_index
                    for dictionary in relative_maximums:
                        ending_height = dictionary.get("height")
                        ending_index = dictionary.get("index")
                        for p in range(relative_starting_index + 1, ending_index):
                            if height[p] < ending_height:
                                volume = ending_height - height[p]
                                total_volume += volume
                        relative_starting_index = ending_index
                
                elif is_bowl:
                    relative_starting_index = bowl_starting_index
                    for dictionary in relative_maximums:
                        ending_height = dictionary.get("height")
                        ending_index = dictionary.get("index")
                        for p in range(relative_starting_index + 1, ending_index):
                            if height[p] < ending_height:
                                volume = ending_height - height[p]
                                total_volume += volume
                        relative_starting_index = ending_index

            elif not is_bowl and height[i] <= height[i + 1]:
                # print("continuing as non-bowl")
                i += 1

            elif not is_bowl and height[i] > height[i + 1]:
                # print("creating new bowl")
                is_bowl = True
                bowl_starting_height = height[i]
                bowl_starting_index = i
                i += 1

            elif is_bowl and height[i] >= bowl_starting_height:
                # print("closing bowl")
                relative_maximums = []
                for n in range(bowl_starting_index + 1, i):
                    volume = bowl_starting_height - height[n]
                    total_volume += volume
                if height[i] > height[i + 1]:
                    bowl_starting_index = i
                    bowl_starting_height = height[i]
                else:
                    is_bowl = False

            elif is_bowl and height[i] > height[i - 1] and height[i] >= height[i + 1]:
                new_rel_maximums = []
                for dictionary in relative_maximums:
                    if dictionary.get("height") >= height[i]:
                        new_rel_maximums.append(dictionary)
                relative_maximums = new_rel_maximums
                relative_maximums.append({"height": height[i], "index": i})

        # print("relative_maximums", relative_maximums)
        # print("total_volume", total_volume)
        return total_volume





solution = Solution()

one = [0, 0, 0]
# result: 0
two = [0, 1, 2]
# result: 0
three = [0, 2, 4, 4, 5]
#result: 0
four = [4, 3, 2, 1]
# result: 0
five = [5, 2, 1, 2]
# result: 1
six = [2, 0, 1, 2]
# result: 3
seven = [0, 0, 0, 2, 0, 1]
# result: 1
eight = [5, 0, 1, 2, 3, 4]
# result: 10
nine = [0,1,0,2,1,0,1,3,2,1,2,1]
# result: 6

ten = [5, 0, 4, 0, 3, 0, 2, 0]
eleven = [5, 0, 1, 0, 2, 0, 3, 0]
twelve = [5, 0, 1, 0, 2, 0, 3, 4]
thirteen = [5, 0, 1, 0, 2, 0, 3, 4, 2, 0, 1]
fourteen = [6,4,2,0,3,2,0,3,1,4,5,3,2,7,5,3,0,1,2,1,3,4,6,8,1,3]
fifteen = [9,6,8,8,5,6,3]
# result: 3

# solution.trap(one)
# solution.trap(two)
# solution.trap(three)
# solution.trap(four)
# solution.trap(five)
# solution.trap(six)
# solution.trap(seven)
# solution.trap(eight)
# solution.trap(nine)
# solution.trap(ten)
# solution.trap(eleven)
# solution.trap(twelve)
# solution.trap(thirteen)
# solution.trap(fourteen)
solution.trap(fifteen)