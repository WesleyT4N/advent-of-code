def has_overlap(a, b):
    return a[0] <= b[1] and b[0] <= a[1]

def get_assignment_overlaps(assignment_file_path):
    overlap_count = 0
    with open(assignment_file_path, "r") as assignment_pairs:
        for pair in assignment_pairs:
            pair = pair.rstrip().split(',')
            range1 = [ int(c) for c in pair[0].split('-') ]
            range2 = [ int(c) for c in pair[1].split('-') ]
            if has_overlap(range1, range2):
                overlap_count += 1
    return overlap_count

print(get_assignment_overlaps("./input"))
