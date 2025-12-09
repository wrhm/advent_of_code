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

assert ((0,1) = earliest_biggest_iv (0,0) [(0,1);(1,1)]);
assert ((1,5) = earliest_biggest_iv (0,0) [(0,1);(1,5);(2,5);(3,3)]);

assert ((1,1) = choose_biggest_after_pos_i_except_for_last_n 0 1 [(0,1);(1,1);(2,1)]);
assert ((1,1) = choose_biggest_after_pos_i_except_for_last_n 0 1 [(0,1);(1,1);(2,1);(3,1)]);
assert ((2,1) = choose_biggest_after_pos_i_except_for_last_n 1 1 [(0,1);(1,1);(2,1);(3,1)]);
assert ((0,8) = choose_biggest_after_pos_i_except_for_last_n (-1) 12 [(0,8);(1,1);(2,1);(3,1);(4,1);(5,1);(6,1);(7,1);(8,1);(9,1);(10,1);(11,1);(12,1);(13,1);(14,9)]);
assert (d03p2 ex_lines03 = 3121910778619);

let char_at i s = 
  (* List.nth (List.of_seq @@ String.to_seq s) i in *)
  String.get s i in

assert ('H' = char_at 0 "Hello");

let width strs = String.length @@ List.nth strs 0 in
let height strs = List.length strs in

let char_from_str_list r c strs = char_at c @@ List.nth strs r in

assert ('e' = char_from_str_list 0 1 ["Hello"]);

let char_from_str_list_or_oob r c strs oob = 
  let h = height strs in
  let w = width strs in
  if r<0 || r>=h || w<0 || w>=c then oob else char_from_str_list r c strs in

assert ('?' = char_from_str_list_or_oob 5 5 [] '?');

let neighbors r c strs nch =
  List.length @@ List.filter
    (fun (ri, ci) -> nch=char_from_str_list_or_oob ri ci strs 'X')
    (List.map (fun (dr,dc) -> (r+dr,c+dc)) [(-1,-1);(-1,0);(-1,1);(0,-1);(0,1);(1,-1);(1,0);(1,1)]) in


let rc_tuples strs = 
  let h = height strs in
  let w = width strs in
  cartesian_product (range 0 (h-1)) (range 0 (w-1)) in

let ex_lines04 = nonempty_lines_from_file (example_dir^"example04.txt") in
print_list print_string "\n" ex_lines04; 
(* print_char @@ char_from_str_list 0 0 ex_lines04; *)
(* print_pair_list print_int print_int @@ rc_tuples ex_lines04; *)
(* print_pair_list print_int print_int @@ List.filter (fun (r,c) ->
'@'=char_from_str_list r c ex_lines04) @@ rc_tuples ex_lines04; *)
let rolls = List.filter (fun (r,c) -> '@'=char_from_str_list r c ex_lines04) @@ rc_tuples ex_lines04 in
let accessible_rolls = List.filter (fun (r,c) -> (neighbors r c ex_lines04 '@')<4) rolls in
print_pair_list print_int print_int accessible_rolls;