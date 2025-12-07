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
