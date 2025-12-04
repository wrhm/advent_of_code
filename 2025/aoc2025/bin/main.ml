(* main.ml

Initialized with: dune init proj aoc2025

To continuously build/run:
* dune build -w
* dune exec aoc2025 -w
*)

let () = 
  Printf.eprintf "\nAnswer 1p1: %d\n" (Aoc2025.Days.d01p1 Aoc2025.Days.lines01);
