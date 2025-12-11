open Io_helpers

let lines01 = nonempty_lines_from_file "inputs/input01.txt"

let parse_letter_number_re =
  Str.regexp "^\\([A-Za-z]+\\)\\([0-9]+\\)$"

let parse_string_to_tuple_option input_string =
  if Str.string_match parse_letter_number_re input_string 0 then
    begin
      let letter_part = Str.matched_group 1 input_string in
      let number_string = Str.matched_group 2 input_string in
      let number_part = int_of_string number_string in
      Some (letter_part, number_part)
    end
  else
    None

let tuple_val  (lr,n) =
  match lr with
  | "L" -> -n
  | _ ->  n

let rec running_modsum acc modv ts =
  match ts with
  | [] -> []
  | ((lr,n)::xs) ->
    let raw = acc + tuple_val (lr,n) in
    let nv = (raw mod modv + modv) mod modv
  in nv::(running_modsum nv modv xs)

let rec running_sum acc ts = 
  match ts with
  | [] -> []
  | ((lr,n)::xs) -> 
    let nv = (acc + tuple_val (lr,n))
  in nv::(running_sum nv xs)

let is100mult x = if (x mod 100 == 0) then 1 else 0

(* how many times is 0 seen when walking distance d (possibly negative) from
pos, excluding the starting point. *)
let rec walked_zeros acc pos d =
  if d == 0 then acc
  else
    let dir = if d > 0 then 1 else -1 in
    let new_pos = pos + dir in
    let new_d = d - dir in
    let add = is100mult new_pos in
    walked_zeros (acc + add) new_pos new_d

(* start at 50. deltas move the dial.
   for each dial movement, walk it slowly to count zero crossings, including
   arrivals at 0. *)
let rec count_zero_crossings4 pos deltas = 
  match deltas with
  | [] -> if (pos mod 100)==0 then 1 else 0
  | (d::ds) -> (walked_zeros 0 pos d) + (count_zero_crossings4 (pos+d) ds)

let as_tuples x = List.filter_map parse_string_to_tuple_option x
let running_modsums x = (running_modsum 50 100 (as_tuples x))

let d01p1 lines = (List.length (List.filter (fun x -> x=0) (running_modsums lines)))
let d01p2 lines =
  count_zero_crossings4 50 (List.map tuple_val (as_tuples lines))

let lines02 = nonempty_lines_from_file "inputs/input02.txt"
let string_halves_as_ints s =
  let n = String.length s in
  let first_len = (n + 1) / 2 in
  let second_len = n - first_len in
  let first = int_of_string @@ String.sub s 0 first_len in
  let second = if second_len = 0 then 0 else int_of_string @@ String.sub s first_len second_len in
  (first, second)

let concat_nums a b = int_of_string @@ (string_of_int a)^(string_of_int b)

let rec range a b =
  if a>b then []
  else a::(range (a+1) b)

let repeats_in_range a b =
  let (a1, a2) = string_halves_as_ints a in
  let (b1, b2) = string_halves_as_ints b in
  let min_y = min (min a1 a2) (min b1 b2) in
  let max_y = max (max a1 a2) (max b1 b2) in
  List.filter (fun x -> (int_of_string a) <= x && x <= (int_of_string b))
  @@ List.map (fun y -> concat_nums y y) @@ range min_y max_y

let list_sum = List.fold_left (+) 0

(* need a fn that repeats a str n times. then need to find all strs up to 10
digits
that are 1-or-more repeats of a digit string. *)

let rec repeat_str n s =
  match n with
  | 0 -> ""
  | 1 -> s
  | _ -> s ^ repeat_str (n-1) s

let cartesian_product list_a list_b =
  List.map (fun a -> List.map (fun b -> (a,b)) list_b) list_a |> List.flatten
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

let str_to_char_list s = List.of_seq @@ String.to_seq s
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

let lines04 = nonempty_lines_from_file "inputs/input04.txt"

let char_at i s = String.get s i
let height strs = List.length strs
let width strs = if height strs = 0 then 0 else String.length @@ List.nth strs 0
let char_from_str_list r c strs = char_at c @@ List.nth strs r

let char_from_str_list_or_oob r c strs oob =
let h = height strs in
let w = width strs in
if r<0 || r>=h || c<0 || c>=w then oob else char_from_str_list r c strs

let neighbors r c strs nch =
  List.filter
    (fun (ri, ci) -> nch=char_from_str_list_or_oob ri ci strs 'X')
    (List.map (fun (dr,dc) -> (r+dr,c+dc)) [(-1,-1);(-1,0);(-1,1);(0,-1);(0,1);(1,-1);(1,0);(1,1)])

let neighbors_among_tuples r c tuples =
  List.filter
    (fun x -> List.mem x tuples)
    (List.map (fun (dr,dc) -> (r+dr,c+dc)) [(-1,-1);(-1,0);(-1,1);(0,-1);(0,1);(1,-1);(1,0);(1,1)])
let rc_tuples strs = 
  let h = height strs in
  let w = width strs in
  cartesian_product (range 0 (h-1)) (range 0 (w-1))

(* let accessible_filter (r,c) = List.length @@ neighbors r c lines '@'<4 *)

let accessible_rolls lines =
  let rolls = List.filter (fun (r,c) -> '@'=char_from_str_list r c lines) @@ rc_tuples lines in
  let ar = List.filter (fun (r,c) -> (List.length @@ neighbors r c lines '@')<4) rolls in
  (rolls, ar)

let accessible_rolls_among_tuples rolls =
  (* let rolls = List.filter (fun (r,c) -> '@'=char_from_str_list r c lines) @@ rc_tuples lines in *)
  let ar = List.filter (fun (r,c) -> (List.length @@ neighbors_among_tuples r c rolls )<4) rolls in
  ar

let d04p1 lines =
  (* List.length @@ accessible_rolls lines *)
  let (_,ar) = accessible_rolls lines in
  (* print_pair_list print_int print_int ar; *)
  List.length ar