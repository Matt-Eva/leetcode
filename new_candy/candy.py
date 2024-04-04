
class Solution:

# start by basically iterating through the array - for loop

# if the child's rating is greater than the previous child, set to higher.

# if the child's rating is less than or equal to the previous child, set to one.
    # in this scenario, we will need to check if the previous child's rating is higher. If it is, then we need to check if it is also getting 1 candy.
    # If it is, then we need to increase the candy it is getting by 1.
    # We then need to repeat this process for that previous post and its own previous post.

# solution timed out:

    # try finding indexes where either side is greater than or equal to its rating.
    # these are our guaranteed 1s.
    # We can think of this problem in terms of valleys and peaks.
    # Valleys are any element that will have a one - either the adjacent elements are equal to it, or the adjacent elements are greater than it.
        # A left valley is a valley whose left side is ascending and right side is equal
        # A right valley is a valley whose right side is ascending and left side is equal
        # A full valley is a valley where both sides are ascending
    # Peaks are any element that is greater than both its neighbors or has one greater than and one equal to.
        # A left peak is a peak whose left side is descending, and right side is equal
        # a right peak is a peak whose right side is descending, and left side is equal
        # A full peak is a peak whose left and right sides are descending
    # Slopes are any element where one side is ascending and another side is descending.
    # We want to first iterate through the array and record valley, peaks, and slopes.

    # we want to record valleys and their orientations in a separate array.
    # We will then proceed through the valley array and use the indexes of those valleys to modify the terrain array.
    # We will also need to check the values from the terrain array to make sure that we are actually updating them correctly.
    # We want to iterate through the terrain array until we reach a peak or a valley.

    def identify_terrain(self, ratings):
        ratings_terrain = []
        valleys = []
        for i in range(len(ratings)):
            current_rating = ratings[i]
            if i == 0:
                next_rating = ratings[i + 1]
                if next_rating > current_rating or next_rating == current_rating:
                    ratings_terrain.append({"terrain_type": "valley", "orientation": "left", "candy_count": 1})
                    valleys.append({"terrain_type": "valley", "orientation": "left", "index": i})
                else:
                    ratings_terrain.append({"terrain_type": "peak", "orietation": "left", "candy_count": 0})
            elif i == len(ratings) - 1:
                prev_rating = ratings[i - 1]
                if prev_rating > current_rating or prev_rating == current_rating:
                    ratings_terrain.append({"terrain_type": "valley", "orientation": "right", "candy_count": 1})
                    valleys.append({"terrain_type": "valley", "orientation": "right", "index": i})
                else:
                    ratings_terrain.append({"terrain_type": "peak","orientation": "left", "candy_count": 0})
            else:
                next_rating = ratings[i + 1]
                prev_rating = ratings[i - 1]
                if next_rating > current_rating and prev_rating > current_rating:
                    ratings_terrain.append({"terrain_type": "valley", 
                    "orientation": "full", "candy_count": 1})
                    valleys.append({"terrain_type": "valley", 
                    "orientation": "full", "index": i})
                elif next_rating > current_rating and prev_rating == current_rating:
                    ratings_terrain.append({"terrain_type": "valley", "orientation": "left", "candy_count": 1})
                    valleys.append({"terrain_type": "valley", "orientation": "left", "index": i})
                elif next_rating == current_rating and prev_rating > current_rating:
                    ratings_terrain.append({"terrain_type": "valley", "orientation": "right", "candy_count": 1})
                    valleys.append({"terrain_type": "valley", "orientation": "right", "index": i})
                elif next_rating == current_rating and prev_rating == current_rating:
                    ratings_terrain.append({"terrain_type": "valley", "orientation": "neutral", "candy_count": 1})
                    valleys.append({"terrain_type": "valley", "orientation": "neutral", "index": i})
                elif next_rating < current_rating and prev_rating < current_rating:
                    ratings_terrain.append({"terrain_type": "peak", "orientation": "full", "candy_count": 0})
                elif next_rating < current_rating and prev_rating == current_rating:
                    ratings_terrain.append({"terrain_type": "peak", "orientation": "left", "candy_count": 0})
                elif next_rating == current_rating and prev_rating < current_rating:
                    ratings_terrain.append({"terrain_type": "peak", "orientation": "right", "candy_count": 0})
                elif next_rating > current_rating and prev_rating < current_rating:
                    ratings_terrain.append({"terrain_type": "slope", "orientation": "left"})
                elif next_rating < current_rating and prev_rating > current_rating:
                    ratings_terrain.append({"terrain_type": "slope", "orientation": "right"})
        return {"terrain": ratings_terrain, "valleys": valleys}

# now that we have our array of valleys and our array with terrain, we need to do the following.
# iterate through the array of valleys.
# For each valley, we need to iterate through the terrain array either right, left, or in both directions, until we reach another valley or peak.
# Unless a valley is neutral, at which point we continue.


    def solution(self, ratings):
        if len(ratings) == 1:
            return 1
        terrain_dict = self.identify_terrain(ratings)
        valleys = terrain_dict.get("valleys")
        print(valleys)
        terrain = terrain_dict.get("terrain")
        for i in range(len(valleys)):
            valley = valleys[i]
            if valley.get("orientation") == "neutral":
                continue
            elif valley.get("orientation") == "left":
                n = valley.get('index') + 1
                slope = True
                while slope:
                    next_terrain = terrain[n]
                    if next_terrain.get("terrain_type") == "valley":
                        slope = False
                    elif next_terrain.get("terrain_type") == "slope":
                        print("slope")
                        terrain[n]["candy_count"] = terrain[n - 1]["candy_count"] + 1
                    elif next_terrain.get("terrain_type") == "peak":
                        new_peak = terrain[n-1]["candy_count"] + 1
                        if new_peak > terrain[n]["candy_count"]:
                            terrain[n]["candy_count"] = new_peak
                        slope = False
                    n += 1
            elif valley.get("orientation") == "right":
                n = valley.get('index') - 1
                slope = True
                while slope:
                    next_terrain = terrain[n]
                    if next_terrain.get("terrain_type") == "valley":
                        slope = False
                    elif next_terrain.get("terrain_type") == "slope":
                        print(slope)
                        terrain[n]["candy_count"] = terrain[n + 1]["candy_count"] + 1
                    elif next_terrain.get("terrain_type") == "peak":
                        new_peak = terrain[n+1]["candy_count"] + 1
                        if new_peak > terrain[n]["candy_count"]:
                            terrain[n]["candy_count"] = new_peak
                        slope = False
                    n -= 1
            elif valley.get("orientation") == "full":
                print(i, valley)
                left = valley.get("index") + 1
                leftSlope = True
                while leftSlope:
                    next_terrain = terrain[left]

                    print(left, valley, next_terrain)
                    if next_terrain.get("terrain_type") == "valley":
                        leftSlope = False
                    elif next_terrain.get("terrain_type") == "slope":
                        terrain[left]["candy_count"] = terrain[left - 1]["candy_count"] + 1
                    elif next_terrain.get("terrain_type") == "peak":
                        new_peak = terrain[left-1]["candy_count"] + 1
                        if new_peak > terrain[left]["candy_count"]:
                            terrain[left]["candy_count"] = new_peak
                        leftSlope = False
                    left += 1
                right = valley.get("index") - 1
                rightSlope = True
                while rightSlope:
                    next_terrain = terrain[right]
                    if next_terrain.get("terrain_type") == "valley":
                        rightSlopelope = False
                    elif next_terrain.get("terrain_type") == "slope":
                        terrain[right]["candy_count"] = terrain[right + 1]["candy_count"] + 1
                    elif next_terrain.get("terrain_type") == "peak":
                        new_peak = terrain[right+1]["candy_count"] + 1
                        if new_peak > terrain[right]["candy_count"]:
                            terrain[right]["candy_count"] = new_peak
                        rightSlope = False
                    right -= 1
        candy_count = 0
        print("terrain", terrain)
        for i in range(len(terrain)):
            candy_count += terrain[i].get("candy_count")
        print(candy_count)
        return candy_count


    # the last solution worked, but it was very expensive. We can do this while maintaining the same array and not using a hash table.
    # we don't actually need to give each child a candy value - we just need to increase the candy count appropriately.
    # We can identify if something is a peak or valley while iterating through it the first time, and allocate a candy count accordingly.
    # First, we need to check if the current rating is a peak, a valley, or a slope.
    # We also need to determine its orientation - left, right, neutral, or full.
    # We do also need to track the distance between peaks and valleys.
        # This is to determine if we need to modify a peak that was previously visited.
        # we need to track distances from peak to valley and valley to peak
        # the distance from a valley to a peak corresponds to the number of candies needed
        # The minimum distance is 1, which is for a valley.
        # We only update the value of a peak if the distance from the peak to the next valley is greater than the distance from the previous valley to that peak.
        # We can keep track of the direction we're heading use a direction variable: -1 is descending from a peak to a valley, 1 is ascending from a valley to a peak, 0 is going from valley to valley.


    def candy(self, ratings):

        if len(ratings) == 1:
                return 1

        candy_count = 0
        valley_to_peak_distance = 1
        peak_to_valley_distance = 1
        direction = 0

        for i in range(len(ratings)):

            if i == 0:
                n = i + 1
                if ratings[n] > ratings[i]:
                    candy_count += 1
                    direction = 1
                elif ratings[n] < ratings[i]:
                    candy_count += 1
                    direction = -1
                else:
                    candy_count += 1
                continue
            
            if i == len(ratings) - 1:
                if direction == 0:
                    candy_count += 1
                elif direction == -1:
                    peak_to_valley_distance += 1
                    candy_count += peak_to_valley_distance
                    if peak_to_valley_distance > valley_to_peak_distance:
                        candy_count += peak_to_valley_distance - valley_to_peak_distance
                else:
                    candy_count += valley_to_peak_distance + 1
                return candy_count
            
            n = i + 1
            q = i - 1

            if ratings[i] > ratings[q] and ratings[i] < ratings[n]:
                direction = 1
                valley_to_peak_distance += 1
                candy_count += valley_to_peak_distance
            if ratings[i] < ratings[q] and ratings[i] > ratings[n]:
                direction = -1
                peak_to_valley_distance += 1
                candy_count += peak_to_valley_distance
            if ratings[i] > ratings[q] and ratings[i] > ratings[n]:
                direction = -1
                valley_to_peak_distance += 1
                candy_count += valley_to_peak_distance









# solution = Solution()
# solution.say_hi()

# ratings1 = [1, 2, 3, 4, 5]
# ratings2 = [5, 4, 3, 2 ,1]
# ratings3 = [2, 1, 3, 4, 2, 5, 6]
# ratings4 = [1, 1, 1, 2, 2, 1, 2, 3]


# solution.candy(ratings1)
# solution.candy(ratings2)
# solution.candy(ratings3)
# solution.candy(ratings4)


