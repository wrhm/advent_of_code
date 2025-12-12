(* main.ml

Initialized with: dune init proj aoc2025

To continuously build/run:
* dune build -w
* dune exec aoc2025 -w
*)

open Aoc2025.Day01
open Aoc2025.Day02
open Aoc2025.Day03
open Aoc2025.Day04
open Aoc2025.Day05
open Aoc2025.Day06

let () = 
  Printf.eprintf "\nAnswer 1p1: %d" @@ d01p1 lines01;
  Printf.eprintf "\nAnswer 1p2: %d" @@ d01p2 lines01;
  Printf.eprintf "\nAnswer 2p1: %d" @@ d02p1 lines02;
  Printf.eprintf "\nAnswer 2p2: %d" @@ d02p2 lines02;
  Printf.eprintf "\nAnswer 3p1: %d" @@ d03p1 lines03;
  Printf.eprintf "\nAnswer 3p2: %d" @@ d03p2 lines03;
  Printf.eprintf "\nAnswer 4p1: %d" @@ d04p1 lines04;
  Printf.eprintf "\nAnswer 4p2: %d" @@ d04p2 lines04;
  Printf.eprintf "\nAnswer 5p1: %d" @@ d05p1 lines05;
  Printf.eprintf "\nAnswer 5p2: %d" @@ d05p2 lines05;
  Printf.eprintf "\nAnswer 6p1: %d" @@ d06p1 lines06;
  Printf.eprintf "\nAnswer 6p2: %d" @@ d06p2 lines06;
  Printf.eprintf "\n"