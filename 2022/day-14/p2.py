SAND_SOURCE = (500, 0)

Coord = tuple[int, int]
Cave = set(Coord)  # cave represented as a set of coords


class SandSimulation:
    def __init__(self, input_file_path: str):
        self.cave, self.y_bounds = self._parse_input(input_file_path)

    def _parse_input(self, input_file_path: str) -> Cave:
        with open(input_file_path) as f:
            rocks = []
            for l in f:
                raw_rock = l.strip().split(" -> ")
                rocks.append(self._parse_rock(raw_rock))

            y_bounds = self._get_y_bounds(rocks)
            return self._build_cave(rocks), y_bounds

    def _parse_rock(self, raw_rock: list[str]) -> list[Coord]:
        rock = []
        for r in raw_rock:
            coord = tuple(int(c) for c in r.split(","))
            rock.append(coord)
        return rock

    def _get_y_bounds(self, rocks: list[list[Coord]]) -> tuple[int, int]:
        max_y = rocks[0][0][1]
        for rock in rocks:
            max_y = max(max_y, *[coord[1] for coord in rock])
        return 0, max_y + 2

    def _build_cave(
        self,
        rocks=list[list[Coord]],
    ) -> Cave:
        cave: Cave = set()
        for rock in rocks:
            cave = self._add_rock(rock, cave)
        return cave

    def _add_rock(self, rock: list[Coord], cave: Cave) -> Cave:
        for i, r in enumerate(rock):
            rx, ry = r
            cave.add((rx, ry))
            if i < len(rock) - 1:
                next_r = rock[i + 1]
                nx, ny = next_r
                dx, dy = nx - rx, ny - ry
                if dx:
                    if nx > rx:
                        for x in range(rx, nx):
                            cave.add((x, ry))
                    else:
                        for x in range(rx, nx - 1, -1):
                            cave.add((x, ry))
                elif dy:
                    if ny > ry:
                        for y in range(ry, ny):
                            cave.add((rx, y))
                    else:
                        for y in range(ry, ny - 1, -1):
                            cave.add((rx, y))
        return cave

    def run_sim(self):
        sand_settled = 0
        last_resting_place = None
        while last_resting_place != SAND_SOURCE:
            last_resting_place = self._resting_place(SAND_SOURCE)
            self.cave.add(last_resting_place)
            sand_settled += 1
        return sand_settled

    def _resting_place(self, init_sand_coord: Coord) -> tuple[int, int] | None:
        cx, cy = init_sand_coord
        while self._above_floor((cx, cy)):
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
        return cx, cy

    def _can_fall_down(self, sand_coord: Coord) -> bool:
        x, y = sand_coord
        return self._above_floor((x, y + 1)) and (x, y + 1) not in self.cave

    def _can_fall_left(self, sand_coord: Coord) -> bool:
        x, y = sand_coord
        return self._above_floor((x - 1, y + 1)) and (x - 1, y + 1) not in self.cave

    def _can_fall_right(self, sand_coord: Coord) -> bool:
        x, y = sand_coord
        return self._above_floor((x + 1, y + 1)) and (x + 1, y + 1) not in self.cave

    def _above_floor(self, sand_coord: Coord) -> bool:
        _, y = sand_coord
        _, max_y = self.y_bounds
        return y < max_y


ss = SandSimulation("./input")
sand_settled = ss.run_sim()
print(sand_settled)
