from collections import deque


class HillClimb:
    def __init__(self, input_file):
        self.start_coords = set()
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
            if c == "E":
                self.end = (row_index, col_index)
                res.append(25)
            elif c == "S" or c == "a":
                self.start_coords.add((row_index, col_index))
                res.append(0)
            else:
                res.append(ord(c) - 97)
        return res

    def fewest_steps_to_end(self):
        # reverse bfs from end to first start coord found
        queue = deque([self.end])
        visited = set([self.end])
        parent_of_coord = {self.end: None}
        while queue:
            current_coord = queue.popleft()
            if current_coord in self.start_coords:
                return self._calc_step_count(parent_of_coord, current_coord)
            for neighb in self.get_neighbors(current_coord, visited):
                parent_of_coord[neighb] = current_coord
                visited.add(neighb)
                queue.append(neighb)

    def get_elevation(self, coord):
        x, y = coord
        return self.map[x][y]

    def get_neighbors(self, coord, visited):
        num_rows = len(self.map)
        num_cols = len(self.map[0])
        x, y = coord
        curr_elevation = self.get_elevation(coord)
        neighbors = []
        for i in [x - 1, x, x + 1]:
            for j in [y - 1, y, y + 1]:
                # within bounds of the map
                neighb = (i, j)
                dist = abs(x - i) + abs(y - j)
                if (
                    dist == 1
                    and 0 <= i < num_rows
                    and 0 <= j < num_cols
                    and curr_elevation - self.get_elevation(neighb) <= 1
                    and neighb not in visited
                ):
                    neighbors.append(neighb)
        return neighbors

    def _calc_step_count(self, parent_of_coord, start):
        curr = start
        step_count = 0
        while curr != self.end:
            curr = parent_of_coord[curr]
            step_count += 1
        return step_count


hc = HillClimb("./input")
print(hc.fewest_steps_to_end())
