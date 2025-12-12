open Io_helpers

let lines03 = nonempty_lines_from_file "inputs/input03.txt"

let rec biggest_except_last m vs =
  match vs with
  | [] -> m
  | [_] -> m
  | (x::y::xs) -> biggest_except_last (max m x) (y::xs)

let rec max_of_list acc vs =
  match vs with
  | [] -> acc
  | [x] -> max acc x
  | (x::xs) -> max_of_list (max acc x) xs

let rec biggest_seen_after dig seen vs =
  match (dig, seen, vs) with
  | (_,_,[]) -> print_string "empty"; failwith "Empty list"
  | (_, true, vs) -> max_of_list 0 vs
  | (_, false, (x::xs)) -> biggest_seen_after dig (dig=x) xs

let digit_char_to_int c = (int_of_char c) - 48
let biggest_joltage s = 
  let int_digits = List.map digit_char_to_int @@ str_to_char_list s in
  let left_digit = biggest_except_last 0 int_digits in
  let right_digit = biggest_seen_after left_digit false int_digits in
  left_digit*10 + right_digit

let d03p1 lines = list_sum @@ List.map biggest_joltage lines

let rec earliest_biggest_iv (acc_i,acc_v) ivs =
  match ivs with
  | [] -> failwith "Empty list"
  | [(i,v)] -> if v>acc_v then (i,v) else (acc_i,acc_v)
  | (i,v)::xs -> if v>acc_v then earliest_biggest_iv (i,v) xs else earliest_biggest_iv (acc_i, acc_v) xs

let choose_biggest_after_pos_i_except_for_last_n i n ordered_pairs =
  let len = List.length ordered_pairs in
  let consider_up_to = len - n + 1 in
  let filtered = List.filter (fun (ind,_) -> ind < consider_up_to) ordered_pairs in
  let after_i = List.filter (fun (ind,_) -> ind > i) filtered in
  earliest_biggest_iv (0,0) after_i

let rec choose_pairs ind n pairs =
  if n = 0 then [] else
  let (i,v) = choose_biggest_after_pos_i_except_for_last_n ind n pairs in
  (i,v)::choose_pairs i (n-1) pairs

(* Choose the subsequence of s that makes the 12 digit number with the highest
possible value. *)
let biggest_joltage_twelve s = 
  let int_digits = List.map digit_char_to_int @@ str_to_char_list s in
  let ordered_pairs = List.combine (range 0 @@ -1 + List.length int_digits) int_digits in
  (* print_string "\n"; *)
  let cpairs = choose_pairs (-1) 12 ordered_pairs in
  let digits = List.map snd cpairs in
  int_of_string (String.concat "" (List.map string_of_int digits))

let d03p2 lines = list_sum @@ List.map biggest_joltage_twelve lines