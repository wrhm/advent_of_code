let read_file filename =
  try
    let ic = In_channel.open_text filename in
    let content = In_channel.input_all ic in
    In_channel.close ic;
    content
  with
  | _ -> "ERROR"

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

let tuple_val  (lr,n) =
  match lr with
  | "L" -> -n
  | _ ->  n

let rec running_sum acc modv ts = 
  match ts with
  | [] -> []
  | ((lr,n)::xs) -> 
    let nv = (acc + tuple_val (lr,n)+modv*999) mod modv
  in nv::(running_sum nv modv xs)

let as_tuples = List.filter_map parse_string_to_tuple_option lines
let running_sums = (running_sum 50 100 (as_tuples))
let d01p1 = (List.length (List.filter (fun x -> x=0) running_sums))
