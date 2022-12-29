import ast
from itertools import groupby

ListOrInt = int | list["ListOrInt"]

PacketList = list[ListOrInt]


class DistressSignal:
    def __init__(self, input_file_path: str):
        self.packet_pairs: list[tuple[PacketList, PacketList]] = self._parse_input(
            input_file_path
        )

    def _parse_input(self, input_file_path: str) -> list[tuple[PacketList, PacketList]]:
        with open(input_file_path) as f:
            lines = [l.strip() for l in f.readlines()]
            grouped_lines = [
                tuple(group)
                for is_non_empty, group in groupby(lines, bool)
                if is_non_empty
            ]
            return [
                (ast.literal_eval(packet_pair[0]), ast.literal_eval(packet_pair[1]))
                for packet_pair in grouped_lines
            ]

    def get_num_packet_pairs_in_right_order(self) -> int:
        packets_in_right_order = [
            self.is_in_right_order(packet[0], packet[1]) for packet in self.packet_pairs
        ]
        sum_of_indices = sum((i + 1) for i, v in enumerate(packets_in_right_order) if v)
        return sum_of_indices

    def is_in_right_order(self, l, r) -> bool | None:
        match l, r:
            case int(l), int(r):
                if l < r:
                    return True
                elif l > r:
                    return False
            case int(l), list(r):
                return self.is_in_right_order([l], r)
            case list(l), int(r):
                return self.is_in_right_order(l, [r])
            case _:
                if l and r:
                    comparison = self.is_in_right_order(l[0], r[0])
                    if type(comparison) == bool:
                        return comparison
                    return self.is_in_right_order(l[1:], r[1:])
                return self.is_in_right_order(len(l), len(r))
        return None


ds = DistressSignal("./input")
print(ds.get_num_packet_pairs_in_right_order())
