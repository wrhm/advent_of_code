open Aoc2025.Days

let () = assert (tuple_val ("",1)==1);

let ex_lines01 = 
  (let all_lines = Str.split (Str.regexp "\n") (read_file "../../../test/example_inputs/example01.txt") in
  List.filter (fun s -> String.length s > 0) all_lines) in
(* Printf.eprintf "%d\n" (List.length ex_lines01);
Printf.eprintf "%s\n" (List.nth ex_lines01 0); *)
assert ((List.length ex_lines01) == 10);
assert (d01p1 ex_lines01 == 3);