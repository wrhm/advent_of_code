(*
$ dune build --watch --terminal-persistence=clear-on-rebuild
$ dune exec ./aoc.exe   
*)
(* let () = Printf.printf "123\n" *)

(* 
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?   
*)

open Base
open Stdio
(* open List *)
(* open Char *)

(* let example = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet"

let lines = String.split_lines example *)

let is_digit x =
  Char.to_int '0' <= Char.to_int x && Char.to_int x <= Char.to_int '9'

let find_digits s =
  let rec h st =
    match st with
    | [] -> []
    | x :: xs -> (if is_digit x then (Char.to_int x)-(Char.to_int '0') :: h xs else h xs)
  in h s

(* let () = printf "%b\n" @@ is_digit '3'; *)
(* let () = printf "%b\n" @@ (find_digits (Base.String.to_list "123")); *)
let a = find_digits (Base.String.to_list "6123def4")
(* let () = List.iter ~f:(printf "%d ") a *)

let () = printf "%d\n" @@ List.length a
let first_elem = match List.hd a with
 | Some x -> x
 | _ -> 0;;
let () = printf "%d\n" @@ first_elem

(* let rec digits s =
  match  *)

(* let solve filename =
  let content = In_channel.read_all filename in
  (* let lines = String.split_lines content in *)
  printf "%s\n" @@ content;
  printf "%s\n" @@ filename;;

let () = solve "01/a.txt" *)

(* let () = printf "%s\n" @@ example; *)