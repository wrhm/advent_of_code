open Io_helpers


let lines04 = nonempty_lines_from_file "inputs/input04.txt"

let char_at i s = String.get s i
let height strs = List.length strs
let width strs = if height strs = 0 then 0 else String.length @@ List.nth strs 0
let char_from_str_list r c strs = char_at c @@ List.nth strs r

let char_from_str_list_or_oob r c strs oob =
let h = height strs in
let w = width strs in
if r<0 || r>=h || c<0 || c>=w then oob else char_from_str_list r c strs

let neighbors r c strs nch =
  List.filter
    (fun (ri, ci) -> nch=char_from_str_list_or_oob ri ci strs 'X')
    (List.map (fun (dr,dc) -> (r+dr,c+dc)) [(-1,-1);(-1,0);(-1,1);(0,-1);(0,1);(1,-1);(1,0);(1,1)])

module IntPair = struct
  type t = int * int
  let compare (x1,y1) (x2,y2) = match compare x1 x2 with 0 -> compare y1 y2 | c -> c
end

module RollSet = Set.Make(IntPair)

let neighbors_among_tuples r c rolls_set =
  List.filter
    (fun x -> RollSet.mem x rolls_set)
    (List.map (fun (dr,dc) -> (r+dr,c+dc)) [(-1,-1);(-1,0);(-1,1);(0,-1);(0,1);(1,-1);(1,0);(1,1)])
let rc_tuples strs =
  let h = height strs in
  let w = width strs in
  cartesian_product (range 0 (h-1)) (range 0 (w-1))

(* let accessible_filter (r,c) = List.length @@ neighbors r c lines '@'<4 *)

let accessible_rolls lines =
  let rolls = List.filter (fun (r,c) -> '@'=char_from_str_list r c lines) @@ rc_tuples lines in
  let ar = List.filter (fun (r,c) -> (List.length @@ neighbors r c lines '@')<4) rolls in
  (rolls, ar)

let accessible_rolls_among_tuples rolls_set =
  let ar = RollSet.filter (fun (r,c) ->
    let neigh = neighbors_among_tuples r c rolls_set in
    (List.length neigh) < 4
  ) rolls_set in
  ar

let d04p1 lines =
  (* List.length @@ accessible_rolls lines *)
  let (_,ar) = accessible_rolls lines in
  (* print_pair_list print_int print_int ar; *)
  List.length ar

let rec remove_rolls_n_times n rs_set =
  if n<=0 then rs_set else
  let ar_set = accessible_rolls_among_tuples rs_set in
  let others_set = RollSet.diff rs_set ar_set in
  (* Printf.eprintf "\n  remaining after removal: %d" (RollSet.cardinal others_set); *)
  flush stderr;
  if RollSet.cardinal ar_set=0 then others_set else
  remove_rolls_n_times (n-1) others_set


let remove_rolls lines =
  flush stderr;
  let (rs_list,_) = accessible_rolls lines in
  let rs_set = RollSet.of_list rs_list in
  let rs_after_set = remove_rolls_n_times 100 rs_set in
  let result = (RollSet.cardinal rs_set) - (RollSet.cardinal rs_after_set) in
  result

let d04p2 lines =
  let result = remove_rolls lines in
  result
