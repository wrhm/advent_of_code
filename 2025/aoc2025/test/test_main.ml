open Aoc2025.Days
open Aoc2025.Io_helpers

let () = assert (tuple_val ("",1)=1);

assert (walked_zeros 0 0 (-5) = 0);
assert (walked_zeros 0 (-5) (5) = 1);
assert (walked_zeros 0 (2) (-3) = 1);

let ex_lines01 = nonempty_lines_from_file "../../../test/example_inputs/example01.txt" in
assert ((List.length ex_lines01) = 10);
assert (d01p1 ex_lines01 = 3);
assert (d01p2 ex_lines01 = 6);

let string_halves_as_ints s =
  let n = String.length s in
  let first_len = (n + 1) / 2 in
  (int_of_string @@ String.sub s 0 first_len, int_of_string @@ String.sub s first_len (n - first_len)) in

(* print_pair print_int print_int @@ string_halves_as_ints "123"; *)
assert (string_halves_as_ints "123" = (12,3));
assert (string_halves_as_ints "1234" = (12,34));
assert (string_halves_as_ints "1012" = (10,12));

let concat_nums a b = int_of_string @@ (string_of_int a)^(string_of_int b) in
assert (concat_nums 12 34 = 1234);

let rec range a b =
  if a>b then []
  else a::(range (a+1) b) in
assert ((range 1 3) = [1;2;3]);

let repeats_in_range a b =
  let (a1, a2) = string_halves_as_ints a in
  let (b1, b2) = string_halves_as_ints b in
  let min_y = min (min a1 a2) (min b1 b2) in
  let max_y = max (max a1 a2) (max b1 b2) in
  List.filter (fun x -> (int_of_string a) <= x && x <= (int_of_string b))
  @@ List.map (fun y -> concat_nums y y) @@ range min_y max_y in

(* print_int_list @@ repeats_in_range "11" "22"; *)
assert ([99]=repeats_in_range "95" "115");
assert ([1188511885]=repeats_in_range "1188511880" "1188511890");
print_string "HELLO ";
print_int_list @@ repeats_in_range "998" "1012";
assert ([1010]=repeats_in_range "998" "1012");


let ex_lines02 = nonempty_lines_from_file "../../../test/example_inputs/example02.txt" in
let pairs02 = List.map (String.split_on_char '-') (String.split_on_char ',' @@
List.nth ex_lines02 0) in
let nums2d02 = List.map (fun t -> repeats_in_range (List.nth t 0) (List.nth t 1)) pairs02 in
let list_sum = List.fold_left (+) 0 in

(* let pairs02asints = List.map (fun slist -> List.map int_of_string slist) pairs02 in *)
(* print_string_list ex_lines02; *)
(* print_string (List.nth ex_lines02 0); *)
(* print_string "\n"; *)
(* print_list print_string " " (String.split_on_char ',' @@ List.nth ex_lines02 0); *)
(* print_string "\n";
print_list_of_lists print_string " " pairs02; *)
print_string "\n";
(* print_list_of_lists print_int " " pairs02asints; *)
print_list_of_lists print_int "/" @@ nums2d02;
(* print_string "\n";
print_int @@ List.fold_left (+) 0 @@ List.map (fun t -> repeats_in_range
(List.nth t 0) (List.nth t 1)) pairs02; *)
print_int @@ list_sum @@ List.map list_sum nums2d02;
