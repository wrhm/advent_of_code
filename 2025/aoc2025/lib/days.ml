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
    let raw = acc + tuple_val (lr,n) in
    let nv = (raw mod modv + modv) mod modv
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

let rec count_zero_crossings3 acc ts =
  match ts with
  | [] -> acc
  | [((delta,res),mres)] -> 
    if ((delta>0) && (res>0)) || ((delta<0) && (res<0)) || (mres==0) then
        let nz = abs(delta/100) in
        (Printf.eprintf "\ndelta:%d res:%d mres:%d zs:%d" delta res mres (acc+nz);
        (acc+nz))
      else acc
  | (((_,pres),_)::((delta,res),mres)::xs) -> 
      let nz = abs((res/100)-(pres/100)) in
        (Printf.eprintf "\ndelta:%d prev:%d res:%d mres:%d zs:%d" delta pres res mres (acc+nz);
        (count_zero_crossings3 (acc+nz) (((delta,res),mres)::xs)))

let is100mult x = if (x mod 100 == 0) then 1 else 0

(* how many times is 0 seen when walking distance d (possibly negative) from
pos, excluding the starting point. *)
let rec walked_zeros acc pos d =
  if d == 0 then acc
  else
    let dir = if d > 0 then 1 else -1 in
    let new_pos = pos + dir in
    let new_d = d - dir in
    let add = is100mult new_pos in
    walked_zeros (acc + add) new_pos new_d

(* start at 50. deltas move the dial.
   for each dial movement, walk it slowly to count zero crossings, including
   arrivals at 0. *)
let rec count_zero_crossings4 pos deltas = 
  match deltas with
  | [] -> if (pos mod 100)==0 then 1 else 0
  | (d::ds) -> (Printf.eprintf "w(%d -> %d -> %d) cz %d\n" pos d (pos+d) (walked_zeros 0 pos d);
  (walked_zeros 0 pos d) + (count_zero_crossings4 (pos+d) ds))

let as_tuples x = List.filter_map parse_string_to_tuple_option x
let running_modsums x = (running_modsum 50 100 (as_tuples x))
let running_sums x = (running_sum 50 (as_tuples x))

(* instructions for d01p1:

For example, suppose the attached document contained the following rotations:

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
Following these rotations would cause the dial to move as follows:

The dial starts by pointing at 50.
The dial is rotated L68 to point at 82.
The dial is rotated L30 to point at 52.
The dial is rotated R48 to point at 0.
The dial is rotated L5 to point at 95.
The dial is rotated R60 to point at 55.
The dial is rotated L55 to point at 0.
The dial is rotated L1 to point at 99.
The dial is rotated L99 to point at 0.
The dial is rotated R14 to point at 14.
The dial is rotated L82 to point at 32.
Because the dial points at 0 a total of three times during this process, the password in this example is 3. *)
let d01p1 lines = (List.length (List.filter (fun x -> x=0) (running_modsums lines)))

(* instructions for d01p2:

you're actually supposed to count the number of times any click causes the dial to point at 0, regardless of whether it happens during a rotation or at the end of one.

Following the same rotations as in the above example, the dial points at zero a few extra times during its rotations:

The dial starts by pointing at 50.
The dial is rotated L68 to point at 82; during this rotation, it points at 0 once.
The dial is rotated L30 to point at 52.
The dial is rotated R48 to point at 0.
The dial is rotated L5 to point at 95.
The dial is rotated R60 to point at 55; during this rotation, it points at 0 once.
The dial is rotated L55 to point at 0.
The dial is rotated L1 to point at 99.
The dial is rotated L99 to point at 0.
The dial is rotated R14 to point at 14.
The dial is rotated L82 to point at 32; during this rotation, it points at 0 once.
In this example, the dial points at 0 three times at the end of a rotation, plus
three more times during a rotation. So, in this example, the new password would
be 6.
Be careful: if the dial were pointing at 50, a single rotation like R1000 would
cause the dial to point at 0 ten times before returning back to 50!
*)
let d01p2 lines =
  (* let fifty_with_orig = (50::(List.map tuple_val (as_tuples lines))) in
  let rs = running_sums ("R0"::lines) in
  (* let rs = running_modsums (ex_lines01) in *)
  let cmb = List.combine fifty_with_orig rs in
  let cmb3 = List.combine cmb (running_modsums ("R0"::lines)) in *)
  (* let czc = count_zero_crossings4 0 cmb3 in *)
  count_zero_crossings4 50 (List.map tuple_val (as_tuples lines))

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
