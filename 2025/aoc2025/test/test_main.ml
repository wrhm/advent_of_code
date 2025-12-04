open Aoc2025.Days

let () = assert (tuple_val ("",1)==1);

let example_input = {ex|
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
|ex} in
let ex_lines = Str.split (Str.regexp "\n") example_input in
assert ((List.length ex_lines) == 10);
assert (d01p1 ex_lines == 3);