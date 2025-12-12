open Io_helpers

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

let lines06 = nonempty_lines_from_file "inputs/input06.txt"