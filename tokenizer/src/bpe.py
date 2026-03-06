type BytePair = tuple[int, int]
type Token = int
type SeenCount = int
type Vocabulary = dict[Token, BytePair]
type CompressedData = tuple[list[int], Vocabulary]


class ClassicBPE:
    """
    In the classic '90s Byte-Pair Encoding algorithm the goal is to maximize data compression.

    Algorithm (Compression):
    1. Find the most frequent adjacent pair (must appear more than once)
    2. Replace all occurrences with a new unique token (starting at 256)
    3. Repeat until no pair appears more than once

    Insights:
    - Most common pair changes post its replacement.
    - Input bytes are converted to list[int] internally to support tokens > 255
    - A pair must repeat more than once to be eligible for replacement. Without this logic, each pair is eligible for replacement.
    - Decompression must process tokens left-to-right; a token may expand into a pair where the left element is itself a token, so processing order determines the final byte order of the decompressed output

    Rabbit Holes:
    - There are ways to optimize the most_common_pair method, including others as well.
    - decompress method can be optimized by using a deque.
    - Compressed output is list[int] + Vocabulary; serializing to disk (endianness, vocab format, versioning) is left to the caller.
    """

    def __init__(self, raw_data: bytes):
        self._raw_data = raw_data
        self._compressed_data = list(raw_data)
        self._next_token: Token = 256
        self._vocab: Vocabulary = {}

    def compression_ratio(self) -> str:
        ratio = len(self._compressed_data) / len(self._raw_data)
        return f"Compressed By: {round(ratio * 100)}%"

    def compress(self) -> CompressedData:
        while True:
            most_common = ClassicBPE._most_common_pair(self._compressed_data)

            if not most_common:
                break

            self._compressed_data = ClassicBPE._replace_all_pairs(
                self._compressed_data, most_common, self._next_token
            )

            self._vocab[self._next_token] = most_common
            self._next_token += 1

        return self._compressed_data, self._vocab

    def decompress(self) -> bytes:
        decompressed_data = bytearray()
        token_stack: list[int] = self._compressed_data.copy()

        # Reverse so that pop() processes tokens left-to-right; without this,
        # the stack processes tokens in reverse order, producing reversed output
        token_stack.reverse()

        while len(token_stack):
            val = token_stack.pop()

            if val < 256:
                decompressed_data.append(val)
            else:
                byte_pair: BytePair = self._vocab[val]
                left, right = byte_pair

                # The leftmost element of the pair must get processed first.
                token_stack.append(right)
                token_stack.append(left)

        return bytes(decompressed_data)

    @staticmethod
    def _most_common_pair(data: list[int]) -> BytePair | None:
        pairs: dict[BytePair, SeenCount] = {}
        most_common: tuple[BytePair | None, SeenCount] = (None, 1)

        for pair in zip(data, data[1:]):
            pairs[pair] = pairs.get(pair, 0) + 1
            seen_count = pairs[pair]

            if most_common[1] < seen_count:
                most_common = (pair, seen_count)

        return most_common[0]

    @staticmethod
    def _replace_all_pairs(
        data: list[int], pair_to_replace: BytePair, replace_by: Token
    ) -> list[int]:
        new_data: list[int] = []
        idx = 0

        while idx < len(data):
            if (
                idx + 1 < len(data)
                and data[idx] == pair_to_replace[0]
                and data[idx + 1] == pair_to_replace[1]
            ):
                new_data.append(replace_by)
                idx += 2
            else:
                new_data.append(data[idx])
                idx += 1

        return new_data


if __name__ == "__main__":
    file = b"aaabdaaabac"

    classic_bpe = ClassicBPE(raw_data=file)

    print(classic_bpe.compress())
    print(classic_bpe.compression_ratio())
    print((classic_bpe.decompress()))
    assert file == classic_bpe.decompress()
