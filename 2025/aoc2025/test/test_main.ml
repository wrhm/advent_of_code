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

let rec earliest_biggest_iv (acc_i,acc_v) ivs =
  match ivs with
  | [] -> failwith "Empty list"
  | [(i,v)] -> if v>acc_v then (i,v) else (acc_i,acc_v)
  | (i,v)::xs -> if v>acc_v then earliest_biggest_iv (i,v) xs else earliest_biggest_iv (acc_i, acc_v) xs in

(* print_pair print_int print_int @@ earliest_biggest_iv (0,0) [(0,1);(1,5);(2,5);(3,3)]; *)
assert ((0,1) = earliest_biggest_iv (0,0) [(0,1);(1,1)]);
assert ((1,5) = earliest_biggest_iv (0,0) [(0,1);(1,5);(2,5);(3,3)]);

let choose_biggest_after_pos_i_except_for_last_n i n ordered_pairs =
  let len = List.length ordered_pairs in
  let consider_up_to = len - n + 1 in
  let filtered = List.filter (fun (ind,_) -> ind < consider_up_to) ordered_pairs in
  let after_i = List.filter (fun (ind,_) -> ind > i) filtered in
  earliest_biggest_iv (0,0) after_i in

assert ((1,1) = choose_biggest_after_pos_i_except_for_last_n 0 1 [(0,1);(1,1);(2,1)]);
assert ((1,1) = choose_biggest_after_pos_i_except_for_last_n 0 1 [(0,1);(1,1);(2,1);(3,1)]);
assert ((2,1) = choose_biggest_after_pos_i_except_for_last_n 1 1 [(0,1);(1,1);(2,1);(3,1)]);
print_pair print_int print_int @@ choose_biggest_after_pos_i_except_for_last_n (-1) 12 [(0,8);(1,1);(2,1);(3,1);(4,1);(5,1);(6,1);(7,1);(8,1);(9,1);(10,1);(11,1);(12,1);(13,1);(14,9)];
assert ((0,8) = choose_biggest_after_pos_i_except_for_last_n (-1) 12 [(0,8);(1,1);(2,1);(3,1);(4,1);(5,1);(6,1);(7,1);(8,1);(9,1);(10,1);(11,1);(12,1);(13,1);(14,9)]);

let rec choose_pairs ind n pairs =
  if n = 0 then [] else
  let (i,v) = choose_biggest_after_pos_i_except_for_last_n ind n pairs in
  (i,v)::choose_pairs i (n-1) pairs in

(* Choose the subsequence of s that makes the 12 digit number with the highest
possible value. *)
let biggest_joltage_twelve s = 
  print_string "\n";
  let int_digits = List.map digit_char_to_int @@ str_to_char_list s in
  print_int_list int_digits;
  let ordered_pairs = List.combine (range 0 @@ -1 + List.length int_digits) int_digits in
  (* print_string "\n"; *)
  print_pair_list print_int print_int ordered_pairs;
  (* let b = choose_biggest_after_pos_i_except_for_last_n (-1) 12 ordered_pairs in
  print_pair print_int print_int b; *)
  (* print_pair print_int print_int @@ choose_biggest_after_pos_i_except_for_last_n (-1) 12 ordered_pairs;
  print_pair print_int print_int @@ choose_biggest_after_pos_i_except_for_last_n (0) 11 ordered_pairs;
  print_pair print_int print_int @@ choose_biggest_after_pos_i_except_for_last_n
  (1) 10 ordered_pairs; *)
  let cpairs = choose_pairs (-1) 12 ordered_pairs in
  print_string "\n";
  let _ = List.map (fun x -> print_pair print_int print_int x; print_string " ") cpairs in
  print_string "\n";
  let _ = List.map (fun (_,v) -> print_int v; print_string " ") cpairs in
  print_string "\n";
  let digits = List.map snd cpairs in
  int_of_string (String.concat "" (List.map string_of_int digits)) in

(* let _ = List.map (fun line -> print_string line; print_string "\n")
ex_lines03 *)
(* print_list_of_lists print_int " " @@ ex_lines03; *)
print_int_list @@ List.map biggest_joltage_twelve ex_lines03;
print_string "\n";
print_int @@ list_sum @@ List.map biggest_joltage_twelve ex_lines03;