max_cal = 0
with open("./input.txt", "r") as input:
    elf_cal = 0
    for line in input:
        line = line.rstrip()
        if line != "":
            elf_cal += int(line)
        else:
            if elf_cal > max_cal:
                max_cal = elf_cal
            elf_cal = 0

print(max_cal)

