
let split s exp = Str.split (Str.regexp exp) s;;

let hash l =
    let rec aux ll tmp res = match ll with
        | [] -> res+tmp
        | ","::t -> aux t 0 (res+tmp)
        | a::t -> let v = Char.code a.[0] in let ntmp = (17*(tmp+v)) mod 256 in aux t ntmp res
    in aux l 0 0
;;


let pars file_name =
  try
    let channel = open_in file_name in
    let line = input_line channel in
    let p1 = hash (split line "") in p1, 0
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;

let a, b = (pars "inputs/input_test15.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input15.txt") in (print_int a; print_string " "; print_int b; print_endline "");;
