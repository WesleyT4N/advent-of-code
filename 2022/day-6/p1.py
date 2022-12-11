def num_chars_before_marker(input_path):
    with open(input_path, "r") as f:
        datastream = f.readline().rstrip()
        l = 0
        r = 4

        seen_chars = set(datastream[l:r])
        while len(seen_chars) < 4 and r < len(datastream):
            l += 1
            r += 1
            seen_chars = set(datastream[l:r])

        return r

print(num_chars_before_marker("./input"))
