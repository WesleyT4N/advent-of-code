def _parse_instr(line):
    line = line.rstrip().split()
    if line[0] == "noop":
        num_cycles = 1
        x_inc = 0
    else:
        num_cycles = 2
        x_inc = int(line[1])
    return num_cycles, x_inc


def _is_sprite_visible(x, cycle):
    cycle_pos = (cycle - 1) % 40
    is_visible = x - 1 <= cycle_pos <= x + 1
    return is_visible


def draw_crt(input_file):
    with open(input_file) as f:
        crt_output = ""
        x = 1
        cycle = 1
        for line in f:
            num_cycles, x_inc = _parse_instr(line)
            for _ in range(num_cycles):
                if _is_sprite_visible(x, cycle):
                    crt_output += "#"
                else:
                    crt_output += "."
                if cycle % 40 == 0:
                    crt_output += "\n"
                cycle += 1
            # increment x after number of cycles from instr
            x += x_inc
        return crt_output


print(draw_crt("./input"))
