let read_file filename =
  try
    let ic = In_channel.open_text filename in
    let content = In_channel.input_all ic in
    In_channel.close ic;
    content
  with
  | _ -> "ERROR: could not read file: "^filename

let nonempty_lines_from_file filename = 
  let all_lines = Str.split (Str.regexp "\n") (read_file filename) in
  List.filter (fun s -> String.length s > 0) all_lines

let print_list item_printer sep lst =
  List.iter (fun x -> item_printer x; print_string sep) lst;
  print_newline ()

let print_list_of_lists item_printer sep lst =
  List.iter (fun x -> print_list item_printer sep x) lst

let print_int_list = print_list print_int " "
let print_string_list = print_list print_string "\n"

let print_pair p1 p2 (x,y) = print_string "("; p1 x; print_string ","; p2 y; print_string ") "

let print_pair_list p1 p2 lst =
  List.iter (fun (x,y) -> print_string "("; p1 x; print_string ","; p2 y; print_string ") ") lst;
  print_newline ()

let print_int_pair_list = print_pair_list print_int print_int

let list_sum = List.fold_left (+) 0

let rec range a b =
  if a>b then []
  else a::(range (a+1) b)

let cartesian_product list_a list_b =
  List.map (fun a -> List.map (fun b -> (a,b)) list_b) list_a |> List.flatten

let rec apply_n n f x = if n=0 then x else apply_n (n-1) f x

let str_to_char_list s = List.of_seq @@ String.to_seq s