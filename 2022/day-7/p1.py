from collections import defaultdict


class Filesystem():
    def __init__(self, input_file_path):
        self.directory_sizes = self._parse_sh_history(input_file_path)

    def _parse_sh_history(self, input_file_path):
        root_dir = {}
        current_path = [("/", root_dir)]
        curr_dir = root_dir
        directory_sizes = defaultdict(int)

        with open(input_file_path) as f:
            raw_lines = [l.rstrip() for l in f.readlines()]

            i = 0
            while i < len(raw_lines):
                line = raw_lines[i]
                if line.startswith("$"):
                    parts = line.split()
                    command = parts[1]

                    if command == "cd":
                        dir = parts[2]
                        if dir == "/":
                            curr_dir = root_dir
                        elif dir == "..":
                            curr_dir = current_path.pop()[1]
                        else:
                            current_path.append((dir, curr_dir))
                            curr_dir = curr_dir[dir]
                        i += 1

                    elif command == "ls":
                        j = i+1
                        while j < len(raw_lines) and not raw_lines[j].startswith("$"):
                            curr_line = raw_lines[j]
                            if curr_line.startswith("dir"):
                                dir_name = curr_line.split()[1]
                                curr_dir[dir_name] = {}
                            else:
                                file_entry = curr_line.split()
                                file_size = int(file_entry[0])
                                file_name = file_entry[1]
                                curr_dir[file_name] = file_size
                                for k in range(1, len(current_path)+1):
                                    directory_sizes['>'.join(
                                        dir_name for dir_name, _ in current_path[:k])] += file_size
                            j += 1
                        i = j

                else:
                    i += 1

        return directory_sizes

    def sum_of_valid_directory_sizes(self):
        return sum([dir_size for dir_size in self.directory_sizes.values() if dir_size <= 100000])

    def sum_of_valid_directory_sizes(self):
        return sum([dir_size for dir_size in self.directory_sizes.values() if dir_size <= 100000])


f = Filesystem("./input")
print(f.sum_of_valid_directory_sizes())
