INIT_SIGNAL_CALC_CYCLE = 20


def _parse_instr(line):
    line = line.rstrip().split()
    if line[0] == "noop":
        num_cycles = 1
        x_inc = 0
    else:
        num_cycles = 2
        x_inc = int(line[1])
    return num_cycles, x_inc


def calc_signal_strengths(input_file):
    with open(input_file) as f:
        signal_strength_sum = 0
        x = 1
        cycle = 1
        curr_signal_strength_calc_cycle = INIT_SIGNAL_CALC_CYCLE

        for line in f:
            num_cycles, x_inc = _parse_instr(line)
            for _ in range(num_cycles):
                if cycle == curr_signal_strength_calc_cycle:
                    signal_strength_sum += cycle * x
                    curr_signal_strength_calc_cycle += 40
                cycle += 1
            # increment x after number of cycles from instr
            x += x_inc
    return signal_strength_sum


print(calc_signal_strengths("./input"))
