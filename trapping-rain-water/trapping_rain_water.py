class Solution:

    # basically, we need to measure the size of the "bowls" contained between different blocks.
    # We then need to add the volume of those bowls together.
    # A bowl only exists if it is contained on either side.
        # Therefore, sections without an opposite side cannot contain rainwater
    # The height of a bowl is dictated by its shortest side.
        # We can identify a bowl once a subsequent side meets or exceeds its height
        # alternatively, we can envision a bowl as a series of individual layered bowls, each with a height of 1.
    # The depth of a bowl varies, and will need to be computed as we move along the bottom of the bowl.
    # we need to compute an initial starting height of a bowl, find the deepest depths of a bowl, and compute the height of where the bowl ends.
    # The ending height of the bowl occurs when the original height is either met, exceeded, or when the height reaches zero.
        # when the height reaches zero, the height of the bowl defaults to the maximum height following the initial height.
        # if the depth is not deeper than this subsequent height, then the bowl holds no water.

    # we are going to need to keep track of certain values

        # a tracker that determines if something meets the criteria for a bowl yet.
            # If the starting height is equal to or less than following adjacent indexes, the starting index moves to the adjacent index, and a bowl has yet to form.
        
        # the starting index
        
        # the starting height

        # the depth layers

        # the width of those depth layers

            # depth layer will be calculated as height.
            # will only be recorded if height is less than starting height
            # if a bowl has not yet been established

            # layers will be stored as a dictionary, whose keys correspond to a height
            # for each key, there will be a starting index and an ending index
            # the difference between the starting and ending index will be to total volume of water held by that layer

        # the ending height

        # the total height

        # the sum of volume of the various depth layers

    # These values can be stored in a hash table (dictionary)

    def prev(self, height):
        # create a variable that tracks total volume
        total_volume = 0
        # create a starting bowl to reset to.
        # layers dict will contain an entry for each layer, which will be a dictionary 
        # need to account for a bowl that has multiple separate instances of the same layer
        # we can iterate through a layer dict and check whether or not the layer dict has an ending index yet
        # if it does, we create a new layer dict.
        # if it doesn't, we set the ending index there.
        # This also means we need to iterate through each layer dict when we end the bowl, and set the ending index to the bowl-ending index - 1.
        # in general, for layer indexes, we start at the index under the first layer block, then end at the index after the final layer block.
        layer = {
            "index": {
                "starting_index": 0,
                "ending_index": None
            }
        }
        reset_bowl_dict = {
            "starting_height": 0,
            "is_bowl": False,
            "layers": {},
        }
        # this dict will be updated for each bowl
        bowl_dict = dict(reset_bowl_dict)

        for i in range(len(height)):
            print(total_volume)
            print(bowl_dict)
            # accounts for index out of range / final entry.
            if i + 1 >= len(height):
                if bowl_dict.get("is_bowl"):
                    if height[i] >= bowl_dict.get("starting_height"):
                        layers = bowl_dict.get("layers")
                        for key in layers:
                            layer = layers.get(key)
                            for index in layer:
                                if not layer.get(index).get("ending_index"):
                                    volume = i - layer.get(index).get("starting_index")
                                    total_volume += volume
                                else:
                                    volume = layer.get(index).get("ending_index") - layer.get(index).get("starting_index")
                                    total_volume += volume
                        return total_volume

                    # if the ending height is not greater that the starting, we need to check and see if there are completed layers
                    # if a layer is complete, we need to add its volume.
                    # if a layer is incomplete, we need to check if the ending height is greater than the height of the layer.
                        # if it is, add the volume
                        # if not, don't add the volume. 
                    else:
                        layers = bowl_dict.get("layers")
                        for key in layers:
                            num_key = int(key)
                            layer = layers.get(key)
                            for index in layer:
                                if not layer.get(index).get("ending_index"):
                                    print("final index", i)
                                    print("final height", height[i])
                                    print("num key", num_key)
                                    print("penultimate volume", total_volume)
                                    if num_key < height[i]:
                                        volume = i - layer.get(index).get("starting_index")
                                        total_volume += volume
                                else:
                                    print("from existing layer", bowl_dict)
                                    volume = layer.get(index).get("ending_index") - layer.get(index).get("starting_index")
                                    total_volume += volume
                        print("final volume", total_volume)
                        return total_volume

                else:
                    return total_volume

            # accounts for instances where no potential bowl exists yet. Moves index forward.
            if not bowl_dict.get("is_bowl") and height[i] <= height[i + 1]:
                continue
            
            # accounts for instances where a potential bowl is created.
            if not bowl_dict.get("is_bowl") and height[i] > height[i + 1]:
                # indicate that we are entering a possible bowl.
                bowl_dict["is_bowl"] = True
                # declare the starting height of the bowl.
                bowl_dict["starting_height"] = height[i]
                # set up a counter to determine how many layers are created
                counter = height[i] - height[i + 1] + 1
                # iterate within counter range to create layers.
                for n in range(1, counter):
                    layer = height[i] - n
                    starting_index = i + 1
                    bowl_dict["layers"][layer] = {starting_index: {"starting_index": starting_index, "ending_index": None}}
                continue
            
            # accounts for instances where a bowl closes.
            if bowl_dict.get("is_bowl") and height[i] >= bowl_dict.get("starting_height"):
                layers = bowl_dict.get("layers")
                for key in layers:
                    layer = layers.get(key)
                    for index in layer:
                        if not layer.get(index).get("ending_index"):
                            volume = i  - layer.get(index).get("starting_index")
                            print("unfinished volume", volume)
                            total_volume += volume
                        else:
                            volume = layer.get(index).get("ending_index") - layer.get(index).get("starting_index")
                            total_volume += volume
                bowl_dict["is_bowl"] = False
                bowl_dict["starting_height"] = 0
                bowl_dict["layers"] = {}
                if not bowl_dict.get("is_bowl") and height[i] > height[i + 1]:
                    # indicate that we are entering a possible bowl.
                    bowl_dict["is_bowl"] = True
                    # declare the starting height of the bowl.
                    bowl_dict["starting_height"] = height[i]
                    # set up a counter to determine how many layers are created
                    counter = height[i] - height[i + 1] + 1
                    # iterate within counter range to create layers.
                    for n in range(1, counter):
                        layer = height[i] - n
                        starting_index = i + 1
                        bowl_dict["layers"][layer] = {starting_index: {"starting_index": starting_index, "ending_index": None}}
                continue

            # accounts for instances where we are within a bowl.
            # in this situation, we will need to manage the creation and closing of new layers.
            # a new layer is created if we descend in height from the previous index.
            # a layer is closed if we ascend in height from the previous index.
            if bowl_dict.get("is_bowl"):
                # accounts for a new layer being created
                if height[i] < height[i - 1]:
                    # need to account if layer dict already exists
                    if not bowl_dict.get("layers").get(f"{height[i]}"):
                        bowl_dict["layers"][height[i]] = {i: {"starting_index": i, "ending_index": None}}
                    else:
                        bowl_dict["layers"][height[i]][i] = {"starting_index": i, "ending_index": None}
                #accounts for a layer being closed.
                elif height[i] > height [i - 1]:
                    layers = bowl_dict.get("layers")
                    for key in layers:
                        num_key = int(key)
                        if num_key < height[i]:
                            layer = layers.get(key)
                            print("layer", layer)
                            for index in layer:
                                if not layer.get(index).get("ending_index"):
                                    bowl_dict["layers"][key][index]["ending_index"] = i
            

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
                print("end index")

            elif not is_bowl and height[i]<= height[i+1]:
                print("still no bowl")

        print(total_volume)
        return total_volume

solution = Solution()

one = [0, 0, 0]
two = [0, 1, 2]
three = [0, 2, 4, 4, 5]

solution.trap(one)