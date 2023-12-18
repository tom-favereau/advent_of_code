let split s exp = Str.split (Str.regexp exp) s;;

let hex_to_int hex_string =
  int_of_string ("0x" ^ hex_string)
;;

let aire lp = 
        let xf, yf = List.hd lp in
        let rec aux lp tmp = match lp with
                | [] -> tmp
                | [a] -> let x, y = a in tmp+(x*yf-xf*y)
                | a::b::tail -> let x1, y1 = a and x2, y2 = b in  
                        aux (b::tail) (tmp+(x1*y2-x2*y1))
        in let a = aux lp 0 in if a < 0 then -a/2 else a/2
;;



let p1 dir dist nb1 lp1 x y = match dir with 
     | "R" -> (nb1+dist), ((x+dist, y)::lp1)
     | "D" -> (nb1+dist), ((x, y-dist)::lp1)
     | "U" -> (nb1+dist), ((x, y+dist)::lp1)     
     | "L" -> (nb1+dist), ((x-dist, y)::lp1)
     | _ -> (print_string dir; failwith "pars p1 matching")
;;

let p2 hex nb2 lp2 x y = 
        let h = String.sub hex 2 ((String.length hex) -4) in 
        let n = String.sub hex ((String.length hex) -2) 1 in
        let dist = hex_to_int h in        
        match n with
        | "0" -> (nb2+dist), ((x+dist, y)::lp2)
        | "1" -> (nb2+dist), ((x, y-dist)::lp2)
        | "2" -> (nb2+dist), ((x-dist, y)::lp2)
        | "3" -> (nb2+dist), ((x, y+dist)::lp2) 
        | _ -> failwith "pars p2 matching"
        
;;

    


let pars file_name =
  try
    let channel = open_in file_name in
    let rec read_lines nb1 lp1 nb2 lp2  =
      try
        let line = input_line channel in
        let arr = split line " " in
        let dir, di, hex = match arr with | a::b::h::[] -> a, b, h | _ -> "", "", "" in
        let dist = int_of_string di in
        let x1, y1 = List.hd lp1 in
        let x2, y2 = List.hd lp2 in
        let nnb1, llp1 = p1 dir dist nb1 lp1 x1 y1 in 
        let nnb2, llp2 = p2 hex nb2 lp2 x2 y2 in
        read_lines nnb1 llp1 nnb2 llp2
      with
      | End_of_file -> (close_in channel; nb1, lp1, nb2, lp2)
    in
    let nb1, lp1, nb2, lp2 = read_lines 1 [(0, 0)] 1 [(0, 0)] in (aire lp1 + nb1/2)+1, (aire lp2 + nb2/2)+1

  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;
                      
let a, b = (pars "inputs/input_test18.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input18.txt") in (print_int a; print_string " "; print_int b; print_endline "");;
