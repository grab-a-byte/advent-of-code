package main

import "core:fmt"
import "core:testing"
import "core:os"
import "core:strconv"
import "core:strings"

main :: proc() {
    file, ok := os.read_entire_file("input.txt", context.allocator)
    if !ok {
        panic("Unable to open file")
    }
    defer delete(file, context.allocator)

    input := make([dynamic]u64)
    it := string(file)
    for line in strings.split_lines_iterator(&it) {
        line_as_int, ok := strconv.parse_u64(line)
        append(&input, line_as_int)
    }

    res := count_increase(input[:])
    fmt.println(res)

}

count_increase :: proc(input: []u64) -> int {
    count := 0
    for i := 1; i<len(input); i += 1 {
        if input[i-1] < input[i] {
            count += 1;
        }
    }

    return count
}

@(test)
test_increasing_numbers :: proc(t: ^testing.T) {
    input := []u64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
    res := count_increase(input)
    testing.expect_value(t, res, 7)
}