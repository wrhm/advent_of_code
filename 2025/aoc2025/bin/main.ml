(* main.ml

Initialized with: dune init proj aoc2025

To continuously build/run:
* dune build -w
* dune exec aoc2025 -w
*)

open Aoc2025.Days

let () = 
  Printf.eprintf "\nAnswer 1p1: %d" @@ d01p1 lines01;
  Printf.eprintf "\nAnswer 1p2: %d" @@ d01p2 lines01;
  Printf.eprintf "\nAnswer 2p1: %d" @@ d02p1 lines02;
  Printf.eprintf "\nAnswer 2p2: %d" @@ d02p2 lines02;
  Printf.eprintf "\nAnswer 3p1: %d" @@ d03p1 lines03;
  Printf.eprintf "\nAnswer 3p2: %d" @@ d03p2 lines03;
  Printf.eprintf "\nAnswer 4p1: %d" @@ d04p1 lines04;
  Printf.eprintf "\nAnswer 4p2: %d" @@ d04p2 lines04;
  Printf.eprintf "\n"