def get_priority(item):
    if item.isupper():
        return ord(item) - 38
    return ord(item) - 96

def sum_priorities(input_path):
    total = 0
    with open(input_path, "r") as rucksacks:
        for rucksack in rucksacks:
            rucksack = rucksack.rstrip()

            compartment_1 = set(rucksack[:len(rucksack)//2])
            compartment_2 = set(rucksack[len(rucksack)//2:])

            intersect = compartment_1 & compartment_2
            total += get_priority(intersect.pop())

    return total

print(sum_priorities('./input'))
