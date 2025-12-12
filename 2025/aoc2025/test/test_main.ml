open Aoc2025.Io_helpers
open Aoc2025.Day01
open Aoc2025.Day02
open Aoc2025.Day03
open Aoc2025.Day04
open Aoc2025.Day05
open Aoc2025.Day06

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

let ex_lines05 = nonempty_lines_from_file (example_dir^"example05.txt") in

assert ((1,2)=nums_from_dashed_pair "1-2");
assert (3 = d05p1 ex_lines05);

assert ([] = coalesce_pairs []);
assert ([(1,2)] = coalesce_pairs [(1,2)]);
assert ([(1,2);(3,4)] = coalesce_pairs [(1,2);(3,4)]);
assert ([(1,4)] = coalesce_pairs [(1,3);(2,4)]);
assert ([(1,5)] = coalesce_pairs [(1,5);(2,4)]);
assert ([(3,5);(10,20)] = coalesce_pairs [(3,5);(10,14);(12,18);(16,20)]);
assert (14 = d05p2 ex_lines05);

let ex_lines06 = nonempty_lines_from_file (example_dir^"example06.txt") in
assert (4277556 = d06p1 ex_lines06);



(* let calculations = List.map List.rev @@ transpose @@ List.map words ex_lines06 in
print_list_of_lists print_string "\n" calculations *)

(* let str_to_char_list *)

let parse_num num_with_spaces =
  let digits = List.filter (fun c -> c <> ' ') @@ str_to_char_list num_with_spaces in
  int_of_string @@ String.of_seq @@ List.to_seq digits in

assert (123 = parse_num "123   ");

let parse_num_and_op s =
  let cl = str_to_char_list s in
  let op = List.nth (List.rev cl) 0 in
  let digits = String.of_seq @@ List.to_seq @@ List.rev @@ List.tl @@ List.rev cl in
  (* print_string ("digits: "^digits); *)
  (parse_num digits,op) in

assert ((123, '*') = parse_num_and_op "123  *");

let rec any f vs =
  match vs with
  | [] -> false
  | (x::xs) -> f x || any f xs in

let string_has_char s c =
  let chs = str_to_char_list s in
  any (fun x -> x=c) chs in

let string_has_digit s =
  let chs = str_to_char_list s in
  any (fun c -> '0' <= c && c <= '9') chs in

assert (true = string_has_digit "123");
assert (false = string_has_digit "abc");

let rec solve_ceph2 strs total stack =
  if List.length strs = 0 then total else
  let s = List.nth strs 0 in
  let ss = List.tl strs in
  if not (string_has_digit s) then solve_ceph2 ss total stack else
  if (string_has_char s '+' || string_has_char s '*') then
    let (n,op) = parse_num_and_op s in
    let res = if op='+' then list_sum (n::stack) else (List.fold_left ( * ) 1 (n::stack)) in
    solve_ceph2 ss (total + res) [] else
  let n = parse_num s in solve_ceph2 ss total (n::stack) in

let charlists = List.map str_to_char_list ex_lines06 in
let numstrs = List.rev @@ List.map (fun cl -> String.of_seq @@ List.to_seq cl) @@ transpose charlists in
let p2 = solve_ceph2 numstrs 0 [] in
(* print_list_of_lists print_char  " " charlists;
print_string "\n=====\n";
print_list_of_lists print_char  " " @@ transpose charlists; *)
print_string "\n=====\n";
print_string_list numstrs;
print_string "\n=====\np2=";
print_int p2