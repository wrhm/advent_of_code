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
(* open Fmt *)
(* open List *)
(* open Char *)

(* let example = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet" *)

(* let lines = String.split_lines example *)

let is_digit x =
  Base.Char.to_int '0' <= Base.Char.to_int x && Base.Char.to_int x <= Base.Char.to_int '9'

let find_digits ch_list =
  let rec h st =
    match st with
    | [] -> []
    | x :: xs -> (if is_digit x then (Char.to_int x)-(Char.to_int '0') :: h xs else h xs)
  in h ch_list

(* let () = printf "%b\n" @@ is_digit '3'; *)
(* let () = printf "%b\n" @@ (find_digits (Base.String.to_list "123")); *)
let a = find_digits (Base.String.to_list "6123def4")
(* let () = List.iter ~f:(printf "%d ") a *)

let () = printf "%d\n" @@ List.length a
let first_elem = match List.hd a with
 | Some x -> x
 | _ -> 0;;
let () = printf "%d\n" @@ first_elem

let print_num n = printf "%d\n" @@ n

let hd_or_default d xs =
  match List.hd xs with
 | Some x -> x
 | _ -> d;; 

let last_or_default d xs =  
  hd_or_default d (List.rev xs);; 

(* 10*first + last *)
let fl_num xs =
  let h = hd_or_default 0 xs in
  let last = last_or_default 0 xs
in h*10+last

let () = print_num ( fl_num [1;2])


(* let formatter_int_list = Fmt.list Fmt.int

let () = Format.asprintf "%a" formatter_int_list [1;2] *)

(* https://gist.github.com/flux77/ab4c20ee28fc3742df29c4f790e6e65f *)
let print_list to_string l =
  let rec loop rem acc = 
    match rem with
    | [] -> acc
    | [x] -> acc ^ (to_string x)
    | (x::xs) -> loop xs (acc ^ (to_string x) ^ "; ") in
  print_string "[";
  print_string (loop l "");
  print_endline "]"

let _print_int_list = print_list Int.to_string

(* let lst = [1;2]
let () = print_int_list lst

let () = print_int_list a *)

(* let rec digits s =
  match  *)

(* let explode_string s = List.init (String.length s) (String.get s);; *)

(* let explode_string *)

let solve filename =
  let content = In_channel.read_all filename in
  let lines = String.split_lines content in
  (* printf "%s\n" @@ content; *)
  (* printf "%s\n" @@ filename;; *)
  let fls = List.map ~f:(fun x -> fl_num (find_digits ( x))) lines in
  (* let fls = List.map (fun x -> x) lines in *)
  _print_int_list fls;;

(* let () = solve "01/a.txt" *)
(* let _x = List.map ~f:(fun x -> x) [1]  *)

(* let () = printf "%s\n" @@ example; *)