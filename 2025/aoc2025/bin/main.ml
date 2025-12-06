(* main.ml

Initialized with: dune init proj aoc2025

To continuously build/run:
* dune build -w
* dune exec aoc2025 -w
*)

let () = 
  Printf.eprintf "\nAnswer 1p1: %d\n" (Aoc2025.Days.d01p1 Aoc2025.Days.lines01);
  Printf.eprintf "\nAnswer 1p2: %d\n" (Aoc2025.Days.d01p2 Aoc2025.Days.lines01);
  Printf.eprintf "\nAnswer 2p1: %d\n" (Aoc2025.Days.d02p1 Aoc2025.Days.lines02);
  (* Printf.eprintf "\nAnswer 2p2: %d\n" (Aoc2025.Days.d02p2 Aoc2025.Days.lines02); *)
