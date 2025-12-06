open Aoc2025.Days

let () = assert (tuple_val ("",1)==1);

assert (walked_zeros 0 0 (-5) == 0);
assert (walked_zeros 0 (-5) (5) == 1);
assert (walked_zeros 0 (2) (-3) == 1);

let ex_lines01 =
  (let all_lines = Str.split (Str.regexp "\n") (read_file "../../../test/example_inputs/example01.txt") in
  List.filter (fun s -> String.length s > 0) all_lines) in
assert ((List.length ex_lines01) == 10);
assert (d01p1 ex_lines01 == 3);
assert (d01p2 ex_lines01 == 6);
