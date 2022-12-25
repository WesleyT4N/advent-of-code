from collections import deque


class HillClimb:
    def __init__(self, input_file):
        self.map = self._parse_map(input_file)

    def _parse_map(self, input_file):
        with open(input_file) as f:
            map = []
            r = 0
            for line in f:
                map.append(self._parse_row(line.strip(), r))
                r += 1
            return map

    def _parse_row(self, row, row_index):
        res = []
        for col_index, c in enumerate(row):
            if c in ["S", "E"]:
                if c == "S":
                    self.start = (row_index, col_index)
                else:
                    self.end = (row_index, col_index)
                res.append(c)
            else:
                res.append(ord(c) - 97)
        return res

    def fewest_steps_to_end(self):
        queue = deque([self.start])
        visited = set()
        parent_of_coord = {self.start: None}
        while queue:
            current_coord = queue.popleft()
            visited.add(current_coord)
            print(current_coord)
            if current_coord == self.end:
                return self._calc_step_count(parent_of_coord)
            for neighb in self.get_neighbors(current_coord):
                if neighb not in visited:
                    parent_of_coord[neighb] = current_coord
                    visited.add(neighb)
                    queue.append(neighb)

    def _calc_step_count(self, parent_of_coord):
        curr = self.end
        step_count = 0
        while curr != self.start:
            curr = parent_of_coord[curr]
            step_count += 1
        return step_count

    def get_elevation(self, coord):
        x, y = coord
        if type(self.map[x][y]) == int:
            return self.map[x][y]
        elif self.map[x][y] == "S":
            return 0
        return 25

    def get_neighbors(self, coord):
        num_rows = len(self.map)
        num_cols = len(self.map[0])
        x, y = coord
        curr_elevation = self.get_elevation(coord)
        neighbors = []
        for i in [x - 1, x, x + 1]:
            for j in [y - 1, y, y + 1]:
                # within bounds of the map
                if (
                    (i, j) != (x, y)
                    and 0 <= i < num_rows
                    and 0 <= j < num_cols
                    and abs(curr_elevation - self.get_elevation((i, j))) <= 1
                ):
                    neighbors.append((i, j))
        return neighbors


hc = HillClimb("./test-input")
print(hc.map)
print(hc.start)
print(hc.fewest_steps_to_end())
