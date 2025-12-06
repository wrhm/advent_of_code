let read_file filename =
  try
    let ic = In_channel.open_text filename in
    let content = In_channel.input_all ic in
    In_channel.close ic;
    content
  with
  | _ -> "ERROR"

let lines01 = 
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

let rec running_modsum acc modv ts = 
  match ts with
  | [] -> []
  | ((lr,n)::xs) -> 
    let nv = (acc + tuple_val (lr,n)+modv*999) mod modv
  in nv::(running_modsum nv modv xs)

let rec running_sum acc ts = 
  match ts with
  | [] -> []
  | ((lr,n)::xs) -> 
    let nv = (acc + tuple_val (lr,n))
  in nv::(running_sum nv xs)

(* ts: [(50,50), (-68,-18), ...] *)
(* neg -> pos: 1 + diff//100 *)
let rec count_zero_crossings acc ts =
  match ts with
  | [] -> acc
  | [(_,_)] -> acc
  | ((_,prev)::(delta,res)::xs) -> 
      if (delta>0) && (prev<0) && (res>0) || (delta<0) && (prev>0) && (res<0) then
        (Printf.eprintf "\ndelta:%d prev:%d res:%d zs:%d" delta prev res (acc+abs(delta)/100);
        (count_zero_crossings (acc+abs(delta)/100) xs))
      else count_zero_crossings acc xs

let rec count_zero_crossings2 acc ts =
  match ts with
  | [] -> acc
  | [((delta,res),mres)] -> 
    if ((delta>0) && (res>0)) || ((delta<0) && (res<0)) || (mres==0) then
        let nz = 1+(abs(delta)/100) in
        (Printf.eprintf "\ndelta:%d res:%d mres:%d zs:%d" delta res mres (acc+nz);
        (acc+nz))
      else acc
  | (((_,pres),_)::((delta,res),mres)::xs) -> 
      if ((delta>0) && (pres<0) && (res>0)) || ((delta<0) && (pres>0) && (res<0)) || (mres==0) then
        let nz = 1+(abs(delta)/100) in
        (Printf.eprintf "\ndelta:%d prev:%d res:%d mres:%d zs:%d" delta pres res mres (acc+nz);
        (count_zero_crossings2 (acc+nz) (((delta,res),mres)::xs)))
      else count_zero_crossings2 acc (((delta,res),mres)::xs)

let as_tuples x = List.filter_map parse_string_to_tuple_option x
let running_modsums x = (running_modsum 50 100 (as_tuples x))
let running_sums x = (running_sum 50 (as_tuples x))
let d01p1 lines = (List.length (List.filter (fun x -> x=0) (running_modsums lines)))
let d01p2 lines =
  let fifty_with_orig = (50::(List.map tuple_val (as_tuples lines))) in
  let rs = running_sums ("R0"::lines) in
  (* let rs = running_modsums (ex_lines01) in *)
  let cmb = List.combine fifty_with_orig rs in
  let cmb3 = List.combine cmb (running_modsums ("R0"::lines)) in
  let czc = count_zero_crossings2 0 cmb3 in
  czc

(* let print_int_list lst =
  List.iter (fun x -> print_int x; print_newline ()) lst *)

let print_list item_printer sep lst =
  List.iter (fun x -> item_printer x; print_string sep) lst;
  print_newline ()

let print_int_list = print_list print_int " "
let print_string_list = print_list print_string " "

let print_pair p1 p2 (x,y) = print_string "("; p1 x; print_string ","; p2 y; print_string ") "

let print_pair_list p1 p2 lst =
  List.iter (fun (x,y) -> print_string "("; p1 x; print_string ","; p2 y; print_string ") ") lst;
  print_newline ()