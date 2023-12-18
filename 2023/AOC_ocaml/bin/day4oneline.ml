let split s exp = Str.split (Str.regexp exp) s;;
let rec pow x n = if n = 0 then 1 else if n mod 2 = 0 then let res = (pow x (n/2)) in res*res else let res = pow x (n/2) in x*res*res;;
let p1 l1 l2 =  let rec aux a = function | [] -> 0 | h::t -> if (h = a) && h <> "" && h <> " " then 1 else aux a t in let rec aux2 = function | [] -> -1 | h::t -> (aux h l2) + (aux2 t) in let nb = aux2 l1 in ( if nb = -1 then 0 else pow 2 nb);;
let pars file_name = try let channel = open_in file_name in let rec read_lines res1 res2  = try let line = input_line channel in let arr = (split (List.hd (List.tl (split line ":"))) "|") in let r1 = p1 (split (List.hd arr) " ") (split (List.hd (List.tl arr)) " ") in read_lines (res1+r1) res2 with | End_of_file -> (close_in channel; res1, res2) in read_lines 0 0 with | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0);;
let a, b = (pars "inputs/input_test04.txt") in (print_int a; print_string " "; print_int b; print_endline "");; let a, b = (pars "inputs/input04.txt") in (print_int a; print_string " "; print_int b; print_endline "");;

