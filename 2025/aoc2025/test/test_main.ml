open Aoc2025.Days

let () = assert (tuple_val ("",1)==1);

let example_input = {ex|
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
|ex} in
let lines = Str.split (Str.regexp "\n") example_input in
(* Printf.eprintf "%x" lines; *)
(* List.iter print_endline lines;
Printf.eprintf "%d" (List.length lines); *)
assert ((List.length lines) == 10);