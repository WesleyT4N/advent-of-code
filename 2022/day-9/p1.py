DIR_TO_VEC = {
    "R": [1, 0],
    "L": [-1, 0],
    "U": [0, 1],
    "D": [0, -1],
}


class RopeSimulation:
    def __init__(self, input_file):
        self.directions = self._parse_input(input_file)

    def adjacent(self, a, b):
        ax, ay = a
        bx, by = b
        return abs(ax - bx) <= 1 and abs(ay - by) <= 1

    def count_tail_visited(self):
        visited = set()
        head = (0, 0)
        tail = (0, 0)
        for direction in self.directions:
            visited.add(tail)

            dx, dy = direction
            hx, hy = head
            head_dest = (hx + dx, hy + dy)

            while head != head_dest:
                head = (
                    int(hx + (dx / abs(dx) if dx else 0)),
                    int(hy + (dy / abs(dy) if dy else 0)),
                )

                hx, hy = head
                if not self.adjacent(head, tail):
                    tx, ty = tail
                    delta_x = (hx - tx) / abs(hx - tx) if hx != tx else 0
                    delta_y = (hy - ty) / abs(hy - ty) if hy != ty else 0
                    tail = (
                        int(tx + delta_x),
                        int(ty + delta_y),
                    )
                    visited.add(tail)

        return len(visited)

    def _parse_input(self, input_file):
        directions = []
        with open(input_file) as f:
            for line in f:
                direction, magnitude = tuple(line.rstrip().split())
                vec = tuple(i * int(magnitude) for i in DIR_TO_VEC[direction])
                directions.append(vec)
        return directions


rs = RopeSimulation("./input")
print("visited", rs.count_tail_visited())
