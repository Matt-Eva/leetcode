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

    def trap(self, height: List[int]) -> int:
        total_volume = 0
        bowl_dict = {
            "starting_height": 0,
            "is_bowl": False
            "layers": {}
            "ending_height": 0
        }

        for i in range(len(height)):

            if i + 1 >= len(height):
                if bowl_dict.get("is_bowl"):
                    if height[i] >= bowl_dict.get("starting_height"):
                        return

                    if height[i] == 0:
                        return

                    if bowl_dict.get("ending_height") == 0:
                        return
                    

                else:
                    return total_volume

            if not bowl_dict.get("is_bowl") and height[i] == height[i + 1]:
                continue
            
            if not bowl_dict.get("is_bowl") and height[i] > height[i + 1]:
                bowl_dict["is_bowl"] = True
                bowl_dict["starting_height"] = height[i]
                bowl_dict["layers"][height[i + 1]] = {starting_index: i, ending_index: i + 1}
                continue
            
            if bowl_dict.get("is_bowl") and height[i] >= bowl_dict.get("starting_height"):
                bowl_dict["is_bowl"] = False
                bowl_dict["ending_height"] = height[i]
                bowl_dict["layers"][bowl_dict.get("starting_he")]
                layers = bowl_dict.get("layers")
                for key in layers:
                    layer = layers.get("key")
                    volume = layer.get("ending_index") - layer_get("starting_index")
                    total_volume += volume
                bowl_dict["layers"] = {}

            if bowl_dict.get("is_bowl") 
            