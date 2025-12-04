let read_file filename =
  try
    let ic = In_channel.open_text filename in
    let content = In_channel.input_all ic in
    In_channel.close ic;
    content
  with
  | _ -> "ERROR"

(* Main execution block *)
(* let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.eprintf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1
  end else
    let filename = Sys.argv.(1) in
    read_file filename *)

(* let example_input = {ex|
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
|ex}
  
let lines = Str.split (Str.regexp "\n") example_input *)

let lines = 
  let all_lines = Str.split (Str.regexp "\n") (read_file "inputs/input01.txt") in
  List.filter (fun s -> String.length s > 0) all_lines

let parse_letter_number_re =
  Str.regexp "^\\([A-Za-z]+\\)\\([0-9]+\\)$"

let parse_string_to_tuple_option input_string =
  if Str.string_match parse_letter_number_re input_string 0 then
    begin
      let letter_part = Str.matched_group 1 input_string in
      let number_string = Str.matched_group 2 input_string in
      let number_part = int_of_string number_string in
      
      Some (letter_part, number_part)
    end
  else
    None

(* let process_and_print input_string =
  let result = parse_string_to_tuple_option input_string in
  
  (* Pattern match the result of the option type *)
  match result with
  | Some (letters, number) -> 
    Printf.printf "Input: \"%s\" -> Success: (\"%s\", %d)\n" 
      input_string letters number
  | None -> 
    Printf.printf "Input: \"%s\" -> Failure: Invalid format (None)\n" 
      input_string *)

(* 1. The List
let numbers = [1; 2; 3; 4; 5];; *)

(* 2. The Accumulator Function (a + element) *)
(* Type: int -> int -> int *)
(* let sum_func accumulator element =
  accumulator + element
  
(* 3. Folding the List *)
let total_sum = 
  List.fold_left 
    sum_func  (* The function to apply at each step *)
    0         (* The initial accumulator value (starting sum) *)
    numbers   (* The list to fold over *)

let () =
  Printf.printf "The list is: [1; 2; 3; 4; 5]\n";
  Printf.printf "The total sum is: %d\n" total_sum *)

let tuple_val  (lr,n) =
  match lr with
  | "L" -> -n
  | _ ->  n

let de_option v =
  match v with
  | Some x ->  x
  | _ ->  failwith ("INVALID")

(* let combine_tuples total (lr,n) =
  match lr with
  | "L" -> total-n
  | _ -> total + n *)

(* let rec running_sum acc ts = 
  match ts with
  | [] -> []
  | ((lr,n)::xs) -> (acc + tuple_val (lr,n))::(running_sum (acc + tuple_val (lr,n)) xs) *)

let rec running_sum acc modv ts = 
  match ts with
  | [] -> []
  | ((lr,n)::xs) -> 
    let nv = (acc + tuple_val (lr,n)+modv*999) mod modv
  in nv::(running_sum nv modv xs)

let as_tuples = List.filter_map parse_string_to_tuple_option lines
let running_sums = (running_sum 50 100 (as_tuples))
let d01p1 = (List.length (List.filter (fun x -> x=0) running_sums))
