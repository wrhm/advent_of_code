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