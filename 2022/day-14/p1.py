import pprint

SAND_SOURCE = (500, 0)

Cave = list[list[bool]]
Coord = tuple[int, int]


class SandSimulation:
    def __init__(self, input_file_path: str):
        self.cave, self.x_bounds, self.y_bounds = self._parse_input(input_file_path)

    def _parse_input(self, input_file_path: str) -> Cave:
        with open(input_file_path) as f:
            rocks = []
            for l in f:
                raw_rock = l.strip().split(" -> ")
                rocks.append(self._parse_rock(raw_rock))

            x_bounds = self._get_x_bounds(rocks)
            y_bounds = self._get_y_bounds(rocks)
            return self._build_cave(x_bounds, y_bounds, rocks), x_bounds, y_bounds

    def _parse_rock(self, raw_rock: list[str]) -> list[Coord]:
        rock = []
        for r in raw_rock:
            coord = tuple(int(c) for c in r.split(","))
            rock.append(coord)
        return rock

    def _get_x_bounds(self, rocks: list[list[Coord]]) -> tuple[int, int]:
        min_x = rocks[0][0][0]
        max_x = 0
        for rock in rocks:
            min_x = min(min_x, *[coord[0] for coord in rock])
            max_x = max(max_x, *[coord[0] for coord in rock])
        return min_x, max_x

    def _get_y_bounds(self, rocks: list[list[Coord]]) -> tuple[int, int]:
        max_y = rocks[0][0][1]
        for rock in rocks:
            max_y = max(max_y, *[coord[1] for coord in rock])
        return 0, max_y

    def _build_cave(
        self,
        x_bounds: tuple[int, int],
        y_bounds: tuple[int, int],
        rocks=list[list[Coord]],
    ) -> Cave:
        min_x, max_x = x_bounds
        _, max_y = y_bounds
        cave: Cave = [
            [False for _ in range(max_x - min_x + 1)] for _ in range(max_y + 1)
        ]
        for rock in rocks:
            cave = self._add_rock(self._scaled_rock(rock, x_bounds), cave)
        return cave

    def _scaled_rock(self, rock: list[Coord], x_bounds: tuple[int, int]):
        min_x, _ = x_bounds
        return [(r[0] - min_x, r[1]) for r in rock]

    def _add_rock(self, rock: list[Coord], cave: Cave) -> Cave:
        for i, r in enumerate(rock):
            if i < len(rock) - 1:
                next_r = rock[i + 1]
            else:
                next_r = None
            rx, ry = r
            if next_r:
                nx, ny = next_r
            else:
                nx, ny = rx, ry

            dx, dy = nx - rx, ny - ry
            if dx:
                if nx > rx:
                    for x in range(rx, nx):
                        cave[ry][x] = True
                else:
                    for x in range(rx, nx - 1, -1):
                        cave[ry][x] = True
            if dy:
                if ny > ry:
                    for y in range(ry, ny):
                        cave[y][rx] = True
                else:
                    for y in range(ry, ny - 1, -1):
                        cave[y][rx] = True
        return cave

    def _print_cave(self):
        cave = self.cave
        output = ""
        for y in range(len(cave)):
            for x in range(len(cave[y])):
                if cave[y][x] is True:
                    output += "#"
                elif cave[y][x] is not False:
                    output += "o"
                else:
                    output += "."
            output += "\n"
        print(output)

    def run_sim(self):
        min_x, _ = self.x_bounds
        sx, sy = SAND_SOURCE[0] - min_x, SAND_SOURCE[1]
        sand_settled = 0
        sand_in_void = False
        while not sand_in_void:
            resting_place = self._resting_place((sx, sy))
            # print(resting_place)
            if resting_place:
                sand_settled += 1
                x, y = resting_place
                self.cave[y][x] = "o"
            else:
                sand_in_void = True
            self._print_cave()
        return sand_settled

    def _resting_place(self, init_sand_coord: Coord) -> tuple[int, int] | None:
        cave = self.cave
        cx, cy = init_sand_coord
        while self._is_in_bounds((cx, cy)):
            if self._can_fall_down((cx, cy)):
                cy += 1
            elif self._can_fall_left((cx, cy)):
                cx -= 1
                cy += 1
            elif self._can_fall_right((cx, cy)):
                cx += 1
                cy += 1
            else:
                return cx, cy
        return None

    def _can_fall_down(self, sand_coord: Coord) -> bool:
        x, y = sand_coord
        return not self._is_in_bounds((x, y + 1)) or self.cave[y + 1][x] is False

    def _can_fall_left(self, sand_coord: Coord) -> bool:
        x, y = sand_coord
        return (
            not self._is_in_bounds((x - 1, y + 1)) or self.cave[y + 1][x - 1] is False
        )

    def _can_fall_right(self, sand_coord: Coord) -> bool:
        x, y = sand_coord
        return (
            not self._is_in_bounds((x + 1, y + 1)) or self.cave[y + 1][x + 1] is False
        )

    def _is_in_bounds(self, sand_coord: Coord) -> bool:
        x, y = sand_coord
        min_x, max_x = self.x_bounds
        min_y, max_y = self.y_bounds
        return (min_x <= x + min_x <= max_x) and (min_y <= y <= max_y)


ss = SandSimulation("./input")
sand_settled = ss.run_sim()
# ss._print_cave()
print(sand_settled)
