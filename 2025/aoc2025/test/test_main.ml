open Aoc2025.Days
open Aoc2025.Io_helpers

let () = assert (tuple_val ("",1)=1);

assert (walked_zeros 0 0 (-5) = 0);
assert (walked_zeros 0 (-5) (5) = 1);
assert (walked_zeros 0 (2) (-3) = 1);

let example_dir = "../../../test/example_inputs/" in

let ex_lines01 = nonempty_lines_from_file (example_dir^"example01.txt") in
assert ((List.length ex_lines01) = 10);
assert (d01p1 ex_lines01 = 3);
assert (d01p2 ex_lines01 = 6);

assert (string_halves_as_ints "123" = (12,3));
assert (string_halves_as_ints "1234" = (12,34));
assert (string_halves_as_ints "1012" = (10,12));

assert (concat_nums 12 34 = 1234);
assert ((range 1 3) = [1;2;3]);
assert ([99]=repeats_in_range "95" "115");
assert ([1188511885]=repeats_in_range "1188511880" "1188511890");
assert ([1010]=repeats_in_range "998" "1012");
assert ([11]=repeats_in_range "2" "18");

assert ("123123"=repeat_str 2 "123");
assert ([(1,4);(1,5);(2,4);(2,5);(3,4);(3,5);]=cartesian_product [1;2;3] [4;5]);

let ex_lines02 = nonempty_lines_from_file (example_dir^"example02.txt") in
assert (d02p1 ex_lines02 = 1227775554);
assert ([1;2;3]=remove_adjacent_repeats [1;1;2;3;3;3]);
assert (d02p2 ex_lines02 = 4174379265);

let ex_lines03 = nonempty_lines_from_file (example_dir^"example03.txt") in
assert (d03p1 ex_lines03 = 357);



