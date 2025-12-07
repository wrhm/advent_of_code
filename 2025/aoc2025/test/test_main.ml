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

print_string "\n";
(* print_string_list repstrs; *)

let ex_lines02 = nonempty_lines_from_file (example_dir^"example02.txt") in
(* let pairs02 = List.map (String.split_on_char '-') (String.split_on_char ',' @@
List.nth ex_lines02 0) in
let nums2d02 = List.map (fun t -> repeats_in_range (List.nth t 0) (List.nth t 1)) pairs02 in
let list_sum = List.fold_left (+) 0 in

print_list_of_lists print_int "/" @@ nums2d02;
print_int @@ list_sum @@ List.map list_sum nums2d02; *)

assert (d02p1 ex_lines02 = 1227775554);

assert ([1;2;3]=remove_adjacent_repeats [1;1;2;3;3;3]);

(* print_list_of_lists print_int " " (d02p2 ex_lines02); *)
(* print_int (d02p2 ex_lines02); *)
print_string "\n";
(* print_int @@ List.length rep_strs; *)
assert (d02p2 ex_lines02 = 4174379265);

let rec biggest_except_last m vs =
  match vs with
  | [] -> m
  | [_] -> m
  | (x::y::xs) -> biggest_except_last (max m x) (y::xs) in

let rec max_of_list acc vs =
  match vs with
  | [] -> acc
  | [x] -> max acc x
  | (x::xs) -> max_of_list (max acc x) xs in

let rec biggest_seen_after dig seen vs =
  match (dig, seen, vs) with
  | (_,_,[]) -> failwith "Empty list"
  | (_, true, vs) -> max_of_list 0 vs
  | (_, false, (x::xs)) -> biggest_seen_after dig (dig=x) xs in

let str_to_char_list s = List.of_seq @@ String.to_seq s in
let digit_char_to_int c = (int_of_char c) - 48 in
let biggest_joltage s = 
  let int_digits = List.map digit_char_to_int @@ str_to_char_list s in
  let left_digit = biggest_except_last 0 int_digits in
  let right_digit = biggest_seen_after left_digit false int_digits in
  left_digit*10 + right_digit in

let ex_lines03 = nonempty_lines_from_file (example_dir^"example03.txt") in
(* let banks = List.map (fun x -> List.map int_of_char x) @@ String.to_list
ex_lines03 *)
let banks = List.map (fun x -> List.map digit_char_to_int x) @@ List.map
str_to_char_list ex_lines03 in
print_list_of_lists print_int " " banks;
let bigjolt = List.map biggest_joltage ex_lines03 in
print_int_list bigjolt;
print_int @@ list_sum bigjolt;



