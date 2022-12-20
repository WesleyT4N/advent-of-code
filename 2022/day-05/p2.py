class Direction():
    def __init__(self, count, src, dest):
        self.count = int(count)
        self.src = int(src)
        self.dest = int(dest)

    def __repr__(self):
        return f"Direction({self.count}, {self.src}, {self.dest})"


class StacksOfCrates():
    def __init__(self, input_file_path):
        stacks, directions = self._parse_file(input_file_path)
        self.stacks = stacks
        self.directions = directions

    def run_directions(self):
        for direction in self.directions:
            chunk = self.stacks[direction.src][-direction.count:]
            self.stacks[direction.src] = self.stacks[direction.src][:-len(chunk)]
            self.stacks[direction.dest] += chunk

    def get_top_of_stacks(self):
        res = ""
        for stack in self.stacks.values():
            if stack:
                res += stack[-1]
        return res

    def _parse_file(self, input_file_path):
        raw_stacks = []
        stack_numbers = ""
        raw_directions = []

        with open(input_file_path) as input_file:
            for line in input_file:
                if "[" in line:
                    raw_stacks.append(line.rstrip())
                elif "move" in line:
                    raw_directions.append(line.rstrip())
                elif "1" in line:
                    stack_numbers = line.rstrip()

        stacks = self._parse_stacks(raw_stacks, stack_numbers)
        directions = self._parse_directions(raw_directions)
        return stacks, directions

    def _get_stack_positions(self, stack_number_line):
        """
        Get map from stack number to index in file string
        """
        stack_positions = {}
        for pos, char in enumerate(stack_number_line):
            if char != " ":
                stack_positions[int(char)] = pos
        return stack_positions

    def _parse_stacks(self, raw_stacks, stack_numbers):
        """
        parse stacks into { <stack_num>: [<stack>] } dict
        """
        stack_positions = self._get_stack_positions(stack_numbers)
        stacks = { num: [] for num in stack_positions}
        for raw_stack in raw_stacks:
            for num in stack_positions:
                position = stack_positions[num]
                if position < len(raw_stack) and raw_stack[position] != " ":
                    stacks[num].append(raw_stack[position])

        return { num: list(reversed(stack)) for num, stack in stacks.items()}

    def _parse_directions(self, raw_directions):
        directions = []
        for dir in raw_directions:
            split_dirs = dir.split(' ')
            count = split_dirs[1]
            src = split_dirs[3]
            dest = split_dirs[5]
            directions.append(Direction(count, src, dest))
        return directions


c = StacksOfCrates("./input")
c.run_directions()
print(c.get_top_of_stacks())
