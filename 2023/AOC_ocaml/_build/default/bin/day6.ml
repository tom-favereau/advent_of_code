
let split s exp = Str.split (Str.regexp exp) s;;

let rec strip s = match s with | [] -> [] | h::t -> if h = " " then strip t else h::(strip t);;

let pars file_name =
  try
    let channel = open_in file_name in
    let time = List.tl (split (input_line channel) " ") in
    let dist = List.tl (split (input_line channel) " ") in
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;
                      
let a, b = (pars "inputs/input_test04.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input04.txt") in (print_int a; print_string " "; print_int b; print_endline "");;

