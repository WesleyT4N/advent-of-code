
class TreeCounter():
    def __init__(self, input_file):
        self.trees = []
        with open(input_file) as tree_map:
            for line in tree_map:
                self.trees.append(list(int(tree_height)
                                  for tree_height in line.rstrip()))

    def count_visible(self):
        visible = 0
        for r, row in enumerate(self.trees):
            for c, tree in enumerate(row):
                if self.is_visible(r, c):
                    visible += 1
        return visible

    def is_on_edge(self, row, col):
        return row in [0, len(self.trees) - 1] or col in [0, len(self.trees[0]) - 1]

    def is_visible(self, row, col):
        if self.is_on_edge(row, col):
            return True

        height = self.trees[row][col]
        row_heights = self.trees[row]
        col_heights = [r[col] for r in self.trees]

        left_heights = row_heights[:col]
        right_heights = row_heights[col + 1:]
        top_heights = col_heights[:row]
        bottom_heights = col_heights[row + 1:]

        return any(
            height > max(heights) for heights in
            (
                left_heights,
                right_heights,
                top_heights,
                bottom_heights
            )
        )


tc = TreeCounter("./input")
print(tc.count_visible())
