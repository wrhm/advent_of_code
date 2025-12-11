open Io_helpers

let lines02 = nonempty_lines_from_file "inputs/input02.txt"
let string_halves_as_ints s =
  let n = String.length s in
  let first_len = (n + 1) / 2 in
  let second_len = n - first_len in
  let first = int_of_string @@ String.sub s 0 first_len in
  let second = if second_len = 0 then 0 else int_of_string @@ String.sub s first_len second_len in
  (first, second)

let concat_nums a b = int_of_string @@ (string_of_int a)^(string_of_int b)

let repeats_in_range a b =
  let (a1, a2) = string_halves_as_ints a in
  let (b1, b2) = string_halves_as_ints b in
  let min_y = min (min a1 a2) (min b1 b2) in
  let max_y = max (max a1 a2) (max b1 b2) in
  List.filter (fun x -> (int_of_string a) <= x && x <= (int_of_string b))
  @@ List.map (fun y -> concat_nums y y) @@ range min_y max_y

(* need a fn that repeats a str n times. then need to find all strs up to 10
digits
that are 1-or-more repeats of a digit string. *)

let rec repeat_str n s =
  match n with
  | 0 -> ""
  | 1 -> s
  | _ -> s ^ repeat_str (n-1) s

let rep_strs = List.filter (fun s -> String.length s <= 10) @@ List.map (fun (a,b) -> repeat_str a (string_of_int b)) @@ cartesian_product (range 2 10) (range 1 99999)

let rec remove_adjacent_repeats list =
  match list with
  | [] -> []
  | [x] -> [x]
  | (x::y::xs) -> if x=y then remove_adjacent_repeats (y::xs) else x::(remove_adjacent_repeats (y::xs))

let rep_strs_deduped = remove_adjacent_repeats @@ List.sort compare rep_strs

let d02p1 lines =
  let pairs02 = List.map (String.split_on_char '-') (String.split_on_char ',' @@ List.nth lines 0) in
  let nums2d02 = List.map (fun t -> repeats_in_range (List.nth t 0) (List.nth t 1)) pairs02 in
  list_sum @@ List.map list_sum nums2d02

let d02p2 lines =
  let pairs02 = List.map (String.split_on_char '-') (String.split_on_char ',' @@ List.nth lines 0) in
  let rep_ints = List.map int_of_string rep_strs_deduped in
  let nums_in_range a b = List.filter (fun x -> a <= x && x <= b) rep_ints in
  let nums_to_sum = List.map (fun t -> nums_in_range (int_of_string @@ List.nth t 0) (int_of_string @@ List.nth t 1)) pairs02 in
  list_sum @@ List.map list_sum nums_to_sum