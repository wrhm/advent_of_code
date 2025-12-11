open Aoc2025.Io_helpers
open Aoc2025.Day01
open Aoc2025.Day02
open Aoc2025.Day03
open Aoc2025.Day04

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

assert ((0,1) = earliest_biggest_iv (0,0) [(0,1);(1,1)]);
assert ((1,5) = earliest_biggest_iv (0,0) [(0,1);(1,5);(2,5);(3,3)]);

assert ((1,1) = choose_biggest_after_pos_i_except_for_last_n 0 1 [(0,1);(1,1);(2,1)]);
assert ((1,1) = choose_biggest_after_pos_i_except_for_last_n 0 1 [(0,1);(1,1);(2,1);(3,1)]);
assert ((2,1) = choose_biggest_after_pos_i_except_for_last_n 1 1 [(0,1);(1,1);(2,1);(3,1)]);
assert ((0,8) = choose_biggest_after_pos_i_except_for_last_n (-1) 12 [(0,8);(1,1);(2,1);(3,1);(4,1);(5,1);(6,1);(7,1);(8,1);(9,1);(10,1);(11,1);(12,1);(13,1);(14,9)]);
assert (d03p2 ex_lines03 = 3121910778619);

assert ('H' = char_at 0 "Hello");
assert ('e' = char_from_str_list 0 1 ["Hello"]);
assert ('?' = char_from_str_list_or_oob 5 5 [] '?');

let ex_lines04 = nonempty_lines_from_file (example_dir^"example04.txt") in
assert (13 = d04p1 ex_lines04);
assert (43=d04p2 ex_lines04);

let num_in_range (low,hi) x = low <= x && x <= hi in

let rec in_any_range rs v =
  match rs with
  | [] -> false
  | (x::xs) -> num_in_range x v || in_any_range xs v in

let nums_from_dashed_pair s =
  let t = String.split_on_char '-' s in
  (int_of_string @@ List.nth t 0, int_of_string @@ List.nth t 1) in

assert ((1,2)=nums_from_dashed_pair "1-2");

let ex_lines05 = nonempty_lines_from_file (example_dir^"example05.txt") in

let d05p1 lines =
  let dashed = List.filter (fun s -> String.contains s '-') lines in
  let nums = List.map int_of_string @@ List.filter (fun s -> not @@ String.contains s '-') lines in
  let ranges = List.map nums_from_dashed_pair dashed in
  let valid_nums = List.filter (fun n -> in_any_range ranges n) nums in
  List.length valid_nums in

(* print_string_list ex_lines05;
print_string "\n"; *)
(* print_int_list valid_nums; *)
assert (3 = d05p1 ex_lines05);
(* print_string "\n";
print_string_list dashed;
print_string "\n";
print_string_list nums; *)