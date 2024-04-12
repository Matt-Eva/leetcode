class Solution:
# Employ concept of a "bowl" that water can fill
# treat that "bowl" as having individual horizontal layers
# verify if a bowl exists, then iterate, calculating the volume of layers

# have a variable that keeps track of total volume
# have a boolean variable that determines if a bowl exists
# have a starting_height variable that determines the starting height of a bowl
# have a variable that is a hash table of layers
# if a bowl does not yet exist, and subsequent levels are greater than or equal to the current level, do nothing and continue forward.
# add layers to the hash table as they appear
# a layer's height is determined by the floor of that layer. 
    #Therefore a layer could have a height of 0
    # This layer would close once the height reached 1 or higher
# each layer has a starting and ending index
    # The starting index is inclusive - the ending index is exclusive
# when the bowl depth drops, create a layer for each level by which the bowl depth has dropped.
# when the bowl depth rises, close any layers whose height is less than the new height.
# when a layer closes, add the volume of the layer to the total-volume variable
    # remove that layer from the hash table
# when the height meets or exceeds the starting height of the bowl, the bowl closes.

    def trap(self, height):
        total_volume = 0
        is_bowl = False
        layers = {}
        starting_height = 0

        for i in range(len(height)):

            if i + 1 >= len(height):
                if is_bowl and height[i] > height[i - 1]:
                    for key in layers:
                        num_key = int(key)
                        if height[i] > key:
                            starting_index = layers.get(key)
                            volume = i - starting_index
                            total_volume += volume

            elif not is_bowl and height[i]<= height[i+1]:
                continue

            elif not is_bowl and height[i] > height[i + 1]:
                is_bowl = True
                starting_height = height[i]

            elif is_bowl and height[i] < height[i - 1]:
                difference = height[i - 1] - height[i]
                for n in range(1, difference + 1):
                    layer = height[i - 1] - n
                    layers[layer] = i
            
            elif is_bowl and height[i] > height[i - 1]:
                closed_layers = []
                for key in layers:
                    num_key = int(key)
                    if height[i] > key:
                        starting_index = layers.get(key)
                        volume = i - starting_index
                        total_volume += volume
                        closed_layers.append(key)
                for n in range(len(closed_layers)):
                    key = closed_layers[n]
                    del layers[key]
                if height[i] >= starting_height:
                    is_bowl = False
                    starting_height = 0

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

# solution.trap(one)
# solution.trap(two)
# solution.trap(three)
# solution.trap(four)
# solution.trap(five)
solution.trap(six)
solution.trap(seven)
solution.trap(eight)