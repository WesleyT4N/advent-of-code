maxes = [0, 0, 0]
with open("./input.txt", "r") as input:
    elf_cal = 0
    for line in input:
        line = line.rstrip()
        if line != "":
            elf_cal += int(line)
        else:
            for i, m in enumerate(maxes):
                if elf_cal > m:
                    maxes[i] = elf_cal
                    break
            elf_cal = 0

print(sum(maxes))
