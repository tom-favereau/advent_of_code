let split s exp = Str.split (Str.regexp exp) s;;

let hash l =
    let rec aux ll tmp res = match ll with
        | [] -> res+tmp
        | ","::t -> aux t 0 (res+tmp)
        | a::t -> let v = Char.code a.[0] in let ntmp = (17*(tmp+v)) mod 256 in aux t ntmp res
    in aux l 0 0
;;

let solve seed inter = 
        let rec min tmp = function | [] -> tmp | (m, _)::t -> if m < tmp then min m t else min tmp t in
        let rec new_inter (d, s, r) tmp = function | [] -> tmp | (a, b)::t ->
                if a > s+r then 0
                else if b < s then 0
                else if a < s && b < s+r then 0
                else if a > s && b > s+r then 0
                else if a < s && b > s+r then 0
                else if a > s && b b < s+r then 0
        in
        let rec through_inter cur li = match li with 
                | [] -> min cur 
                | (a, b)::t -> (let rec subaux tmp = function )

let pars file_name =
  try
    let channel inter tmp = open_in file_name in
    let rec read_lines tmp  =
      try
        let line = input_line channel in
        if line = "" then (input_line channel; read_lines (tmp::inter) [])
        else let i = match split line " " with | a::b::c::_ -> a, b, c | _ -> 0, 0, 0 in
        read_lines inter i::tmp
      with
      | End_of_file -> (close_in channel; inter)
    in
    let pars_seed = input_line channel in let inter = read_lines [] [] in
    let rec seed1 tmp = function | [] -> tmp | h::t -> seed1 (h, h)::tmp t in
    let rec seed2 tmp = function | [] -> tmp | s::r::t -> seed2 (s, s+r)::tmp t in
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;


let a, b = (pars "inputs/input_test15.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input15.txt") in (print_int a; print_string " "; print_int b; print_endline "");;
