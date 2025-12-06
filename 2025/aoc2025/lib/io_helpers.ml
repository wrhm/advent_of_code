let read_file filename =
  try
    let ic = In_channel.open_text filename in
    let content = In_channel.input_all ic in
    In_channel.close ic;
    content
  with
  | _ -> "ERROR"

let print_list item_printer sep lst =
  List.iter (fun x -> item_printer x; print_string sep) lst;
  print_newline ()

let print_int_list = print_list print_int " "
let print_string_list = print_list print_string " "

let print_pair p1 p2 (x,y) = print_string "("; p1 x; print_string ","; p2 y; print_string ") "

let print_pair_list p1 p2 lst =
  List.iter (fun (x,y) -> print_string "("; p1 x; print_string ","; p2 y; print_string ") ") lst;
  print_newline ()
