def get_priority(item):
    if item.isupper():
        return ord(item) - 38
    return ord(item) - 96

def get_group_priority(group):
    intersect = set(group[0]) & set(group[1]) & set(group[2])
    if intersect:
        return get_priority(intersect.pop())
    return 0

def sum_priorities(input_path):
    total = 0
    with open(input_path, "r") as rucksacks:
        group = ["", "", ""]
        for i, rucksack in enumerate(rucksacks):
            rucksack = rucksack.rstrip()
            group[i % 3] = rucksack
            if i != 0 and i % 3 == 2:
                total += get_group_priority(group)
    return total

print(sum_priorities('./input'))
