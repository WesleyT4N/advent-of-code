import pprint


class Monkey:
    def __init__(self, raw_monkey_input):
        self.items = self._parse_items(raw_monkey_input[1])
        self.op = self._parse_op(raw_monkey_input[2])
        self.test = self._parse_test(raw_monkey_input[3:])

    def _parse_items(self, raw_items):
        return [int(i) for i in raw_items.replace(",", "").split()[2:]]

    def _parse_op(self, raw_op):
        self.raw_op = raw_op
        op = raw_op[raw_op.index("=") + 1 :].strip()

    def _parse_test(self, raw_test_lines):
        self.raw_test = raw_test_lines
        return []

    def __repr__(self):
        return f"<Monkey items:{self.items}, op:{self.raw_op}, test:{self.raw_test}>"


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


mks = MonkeyBusinessSimulator("./test-input")
pprint.pprint(mks.monkeys)
