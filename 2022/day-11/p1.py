from collections import deque


class Monkey:
    def __init__(self, raw_monkey_input):
        self.items = self._parse_items(raw_monkey_input[1])
        self.op = self._parse_op(raw_monkey_input[2])
        self.test = self._parse_test(raw_monkey_input[3:])

    def _parse_items(self, raw_items):
        return deque(int(i) for i in raw_items.replace(",", "").split()[2:])

    def _parse_op(self, raw_op):
        self.raw_op = raw_op
        op = raw_op[raw_op.index("=") + 1 :].strip()
        return lambda old: eval(op)

    def _parse_test(self, raw_test_lines):
        self.raw_test = raw_test_lines

        divisible_by = int(raw_test_lines[0].split()[-1])
        monkey_num_if_true = int(raw_test_lines[1].split()[-1])
        monkey_num_if_false = int(raw_test_lines[2].split()[-1])

        return (
            lambda val: monkey_num_if_true
            if val % divisible_by == 0
            else monkey_num_if_false
        )

    def __repr__(self):
        return (
            f"<Monkey items:{list(self.items)}, op:{self.raw_op}, test:{self.raw_test}>"
        )


class MonkeyBusinessSimulator:
    def __init__(self, input_file_path):
        self.monkeys = self._parse_input(input_file_path)

    def _parse_input(self, input_file_path):
        with open(input_file_path) as f:
            monkeys = []
            curr_monkey = []
            for line in f:
                curr_line = line.strip()
                if not curr_line:
                    monkeys.append(Monkey(curr_monkey))
                    curr_monkey = []
                else:
                    curr_monkey.append(curr_line)
            monkeys.append(Monkey(curr_monkey))
            return monkeys

    def run(self, num_rounds):
        monkey_inspect_counts = [0 for _ in range(len(self.monkeys))]
        for round in range(num_rounds):
            for monkey_num, monkey in enumerate(self.monkeys):
                while monkey.items:
                    print("Monkey", monkey_num)
                    item = monkey.items.popleft()
                    print("inspecting", item)
                    item = monkey.op(item)
                    monkey_inspect_counts[monkey_num] += 1
                    print("new item worry level", item)
                    item //= 3
                    print("new item worry level after div", item)
                    target_monkey_num = monkey.test(item)
                    print("throwing item", item, "to Monkey", target_monkey_num)
                    self.monkeys[target_monkey_num].items.append(item)
        print("monkey inspections", monkey_inspect_counts)
        print("monkey business", self._calc_monkey_business(monkey_inspect_counts))

    def _calc_monkey_business(self, monkey_inspect_counts):
        sorted_counts = sorted(monkey_inspect_counts)
        return sorted_counts[-2] * sorted_counts[-1]


mks = MonkeyBusinessSimulator("./input")
mks.run(20)
