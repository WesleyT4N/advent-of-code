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
        knots = [(0, 0) for _ in range(10)]
        for direction in self.directions:
            visited.add(knots[-1])

            dx, dy = direction
            hx, hy = knots[0]
            head_dest = (hx + dx, hy + dy)
            while knots[0] != head_dest:
                knots[0] = (
                    int(hx + (dx / abs(dx) if dx else 0)),
                    int(hy + (dy / abs(dy) if dy else 0)),
                )
                hx, hy = knots[0]

                for i in range(1, len(knots)):
                    if not self.adjacent(knots[i - 1], knots[i]):
                        px, py = knots[i - 1]
                        kx, ky = knots[i]
                        delta_x = (px - kx) / abs(px - kx) if (px - kx) else 0
                        delta_y = (py - ky) / abs(py - ky) if (py - ky) else 0
                        knots[i] = (
                            int(kx + delta_x),
                            int(ky + delta_y),
                        )

                visited.add(knots[-1])

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
