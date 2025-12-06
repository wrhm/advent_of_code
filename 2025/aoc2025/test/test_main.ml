open Aoc2025.Days

let () = assert (tuple_val ("",1)==1);

assert (walked_zeros 0 0 (-5) == 0);
assert (walked_zeros 0 (-5) (5) == 1);
Printf.eprintf "wz: %d" (walked_zeros 0 (2) (-3));
assert (walked_zeros 0 (2) (-3) == 1);

let ex_lines01 =
  (let all_lines = Str.split (Str.regexp "\n") (read_file "../../../test/example_inputs/example01.txt") in
  List.filter (fun s -> String.length s > 0) all_lines) in
let fifty_with_orig = (50::(List.map tuple_val (as_tuples ex_lines01))) in
let rs = running_sums ("R0"::ex_lines01) in
(* let rs = running_modsums (ex_lines01) in *)
let cmb = List.combine fifty_with_orig rs in
let cmb3 = List.combine cmb (running_modsums ("R0"::ex_lines01)) in
(* let czc = count_zero_crossings4 0 cmb3 in *)
let czc = count_zero_crossings4 50 (List.map tuple_val (as_tuples ex_lines01)) in
(* Printf.eprintf "%d\n" (List.length ex_lines01);
Printf.eprintf "%s\n" (List.nth ex_lines01 0); *)
assert ((List.length ex_lines01) == 10);
assert (d01p1 ex_lines01 == 3);
(* print_int_list ex_lines01; *)
(* List.iter (fun x -> print_string x; print_string " ") ex_lines01; *)
print_string_list ex_lines01;
(* let fifty_with_orig = (50::(List.map tuple_val (as_tuples ex_lines01))); *)
print_int_list fifty_with_orig;
(* print_list print_int " " (running_sums (50::(List.map tuple_val (as_tuples ex_lines01)))); *)
(* print_list print_int " " (running_sums ex_lines01); *)
print_int_list rs;
print_int_list (running_modsums ("R0"::ex_lines01));
print_pair_list print_int print_int cmb;
print_pair_list (fun t -> print_pair print_int print_int t) print_int cmb3;
(* print_int (count_zero_crossings 0 cmb); *)
print_int czc;
print_int (d01p2 ex_lines01);
assert (d01p2 ex_lines01 == 6);
