let split s exp = Str.split (Str.regexp exp) s;;

let rec pow x n = if n = 0 then 1 else if n mod 2 = 0 then let res = (pow x (n/2)) in res*res else let res = pow x (n/2) in x*res*res;;

let number l1 l2 =  
        let rec aux a = function 
                | [] -> 0
                | h::t -> if (h = a) && h <> "" && h <> " " then 1 else aux a t
        in let rec aux2 = function | [] -> 0 | h::t -> (aux h l2) + (aux2 t) in
        aux2 l1
;;

let p1 l1 l2 = let nb = number l1 l2 in if nb = 0 then 0 else pow 2 (nb-1);;

let p2 l = let n = List.length l in let arr = Array.make n 1 in let rec aux tmp i = function
        | [] -> tmp
        | [_] -> tmp+arr.(i)
        | h::t -> for k = i+1 to min (i+h) (n-1) do arr.(k) <- arr.(k)+1*arr.(i) done; aux (tmp+arr.(i)) (i+1) t
        in aux 0 0 l
;; 

let pars file_name =
  try
    let channel = open_in file_name in
    let rec read_lines res1 res2  =
      try
        let line = input_line channel in
        let arr = (split (List.hd (List.tl (split line ":"))) "|") in
        let r1 = p1 (split (List.hd arr) " ") (split (List.hd (List.tl arr)) " ") in 
        let r2 = number (split (List.hd arr) " ") (split (List.hd (List.tl arr)) " ") in        
        read_lines (res1+r1) (r2::res2)
      with
      | End_of_file -> (close_in channel; res1, p2 (List.rev res2))
    in
    read_lines 0 []

  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;
                      
let a, b = (pars "inputs/input_test04.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input04.txt") in (print_int a; print_string " "; print_int b; print_endline "");;

