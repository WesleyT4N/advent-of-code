class TreeCounter():
    def __init__(self, input_file):
        self.trees = []
        with open(input_file) as tree_map:
            for line in tree_map:
                self.trees.append(list(int(tree_height)
                                  for tree_height in line.rstrip()))

    def max_scenic_score(self):
        max_scenic_score = 0
        return max(
            self.scenic_score(r, c)
            for r, row in enumerate(self.trees)
            for c, _ in enumerate(row)
        )

    def scenic_score(self, row, col):
        height = self.trees[row][col]
        row_heights = self.trees[row]
        col_heights = [r[col] for r in self.trees]

        left_heights = row_heights[:col]
        left_viewing_distance = self.viewing_distance(
            reversed(left_heights), height)

        right_heights = row_heights[col + 1:]
        right_viewing_distance = self.viewing_distance(right_heights, height)

        top_heights = col_heights[:row]
        top_viewing_distance = self.viewing_distance(
            reversed(top_heights), height)

        bottom_heights = col_heights[row + 1:]
        bottom_viewing_distance = self.viewing_distance(bottom_heights, height)

        return (left_viewing_distance * right_viewing_distance *
                top_viewing_distance * bottom_viewing_distance)

    def viewing_distance(self, trees_in_direction, height):
        viewing_distance = 0
        for h in trees_in_direction:
            viewing_distance += 1
            if h >= height:
                break
        return viewing_distance


tc = TreeCounter("./input")
print(tc.max_scenic_score())
