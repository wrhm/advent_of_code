open Io_helpers

let lines05 = nonempty_lines_from_file "inputs/input05.txt"

let num_in_range (low,hi) x = low <= x && x <= hi

let rec in_any_range rs v =
  match rs with
  | [] -> false
  | (x::xs) -> num_in_range x v || in_any_range xs v

let nums_from_dashed_pair s =
  let t = String.split_on_char '-' s in
  (int_of_string @@ List.nth t 0, int_of_string @@ List.nth t 1)

let d05p1 lines =
  let dashed = List.filter (fun s -> String.contains s '-') lines in
  let nums = List.map int_of_string @@ List.filter (fun s -> not @@ String.contains s '-') lines in
  let ranges = List.map nums_from_dashed_pair dashed in
  let valid_nums = List.filter (fun n -> in_any_range ranges n) nums in
  List.length valid_nums

let compare_by_first (a,_) (b,_) = compare a b

let sort_int_pairs ps = List.sort compare_by_first ps

let rec coalesce_pairs ps =
  match ps with
  | [] -> []
  | [x] -> [x]
  | (a,b)::(c,d)::xs -> 
    if b>=c then coalesce_pairs @@ (a,max b d)::xs else
    (a,b)::(coalesce_pairs @@ (c,d)::xs)

let d05p2 lines =
  let dashed = List.filter (fun s -> String.contains s '-') lines in
  let ranges = List.map nums_from_dashed_pair dashed in
  let srs = sort_int_pairs ranges in
  let co = coalesce_pairs srs in 
  list_sum @@ List.map (fun (a,b) -> b-a+1) co
