def fully_contains(a, b):
    return a[0] <= b[0] <= b[1] <= a[1]

def get_assignment_overlaps(assignment_file_path):
    fully_contains_count = 0
    with open(assignment_file_path, "r") as assignment_pairs:
        for pair in assignment_pairs:
            pair = pair.rstrip().split(',')
            range1 = [ int(c) for c in pair[0].split('-') ]
            range2 = [ int(c) for c in  pair[1].split('-') ]
            if fully_contains(range1, range2) or fully_contains(range2, range1):
                fully_contains_count += 1
    return fully_contains_count

print(get_assignment_overlaps("./input"))
