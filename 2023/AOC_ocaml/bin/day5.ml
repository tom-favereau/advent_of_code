let split s exp = Str.split (Str.regexp exp) s;;

let solve seed inter = 
        let rec min tmp = function | [] -> tmp | (m, _)::t -> if m < tmp then min m t else min tmp t in
        let rec new_inter (d, s, r) tmp = function | [] -> (tmp, []) | (a, b)::t ->
                if a > s+r then 
                  let ni, pri = new_inter (d, s, r) tmp t in ni, ((a, b)::pri) 
                else if b < s then 
                  let ni, pri = new_inter (d, s, r) tmp t in ni, ((a, b)::pri) 
                else if a <= s && b <= s+r then 
                  let ni, pri = new_inter (d, s, r) ((d, d+b-s)::tmp) t in ni, ((a, s-1)::pri)
                else if a >= s && b >= s+r then 
                  let ni, pri = new_inter (d, s, r) ((d+a-s, d+r)::tmp) t in ni, ((s+r, b)::pri)
                else if a <= s && b >= s+r then 
                  let ni, pri = new_inter (d, s, r) ((d, d+r)::tmp) t in ni, ((a, s-1)::(s+r+1, b)::pri)
                else if a > s && b < s+r then 
                  let ni, pri = new_inter (d, s, r) ((d+a-s, s+b-a)::tmp) t in ni, pri
                else [], []
        in
        let rec clean = function | [] -> [] | (a, b)::t -> if a > b then clean t 
                                                        else ((a, b)::(clean t)) 
        in
        let rec through_step cur llrange = match llrange with 
                | [] -> min cur 
                | h::t -> let rec loop_over_range res = function 
                            | [] -> res
                            | a::q -> let ni, pri = new_inter a [] cur in  
;;



let pars file_name =
  try
    let channel inter tmp = open_in file_name in
    let rec read_lines tmp  =
      try
        let line = input_line channel in
        if line = "" then (input_line channel; read_lines ((tmp)::inter) [])
        else let i = match split line " " with | a::b::c::_ -> a, b, c | _ -> 0, 0, 0 in
        read_lines inter i::tmp
      with
      | End_of_file -> (close_in channel; inter)
    in
    let pars_seed = input_line channel in let inter = read_lines [] [] in
    let rec seed1 tmp = function | [] -> tmp | h::t -> seed1 (h, h)::tmp t in
    let rec seed2 tmp = function | [] -> tmp | s::r::t -> seed2 (s, s+r)::tmp t in
    let r1 = solve (seed1 pars_seed) inter and r2 = solve (seed2 pars_seed) inter in
    r1, r2
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;


let a, b = (pars "inputs/input_test15.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input15.txt") in (print_int a; print_string " "; print_int b; print_endline "");;
