open Io_helpers

let lines06 = nonempty_lines_from_file "inputs/input06.txt"

let words s =
  List.filter (fun w -> w <> "") @@ String.split_on_char ' ' s

let rec transpose m =
  match m with
  | [] :: _ -> []
  | [] -> []
  | rows ->
    (List.map List.hd rows)::(transpose (List.map List.tl rows))

let solve_ceph xs =
  let op = List.hd xs in
  let vs = List.map int_of_string @@ List.tl xs in
  if op="+" then List.fold_left (+) 0 vs else
    List.fold_left ( * ) 1 vs

let d06p1 lines =
  let calculations = List.map List.rev @@ transpose @@ List.map words lines in
  list_sum @@ List.map solve_ceph calculations

let parse_num num_with_spaces =
  let digits = List.filter (fun c -> c <> ' ') @@ str_to_char_list num_with_spaces in
  int_of_string @@ String.of_seq @@ List.to_seq digits

let parse_num_and_op s =
  let cl = str_to_char_list s in
  let op = List.nth (List.rev cl) 0 in
  let digits = String.of_seq @@ List.to_seq @@ List.rev @@ List.tl @@ List.rev cl in
  (* print_string ("digits: "^digits); *)
  (parse_num digits,op)

let rec any f vs =
  match vs with
  | [] -> false
  | (x::xs) -> f x || any f xs

let string_has_char s c =
  let chs = str_to_char_list s in
  any (fun x -> x=c) chs

let string_has_digit s =
  let chs = str_to_char_list s in
  any (fun c -> '0' <= c && c <= '9') chs

let rec solve_ceph2 strs total stack =
  if List.length strs = 0 then total else
  let s = List.nth strs 0 in
  let ss = List.tl strs in
  if not (string_has_digit s) then solve_ceph2 ss total stack else
  if (string_has_char s '+' || string_has_char s '*') then
    let (n,op) = parse_num_and_op s in
    let res = if op='+' then list_sum (n::stack) else (List.fold_left ( * ) 1 (n::stack)) in
    solve_ceph2 ss (total + res) [] else
  let n = parse_num s in solve_ceph2 ss total (n::stack)

let d06p2 lines =
  let charlists = List.map str_to_char_list lines in
  let numstrs = List.rev @@ List.map (fun cl -> String.of_seq @@ List.to_seq cl) @@ transpose charlists in
  solve_ceph2 numstrs 0 []
