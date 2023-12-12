(*open Str;;*)

(*let string_to_list str = str |> String.to_seq |> List.of_seq;;*)

let split s exp = Str.split (Str.regexp exp) s;;

(*
let last str =
  let length = String.length str in
  if length > 0 then
    Some str.[length - 1]
  else
    None
;;
*)

let rec doPart1 w num tmp first = match w, num with 
        | [], _ -> if num = [] then 1 else 0
        | w, [] -> let rec aux = function | [] -> true | h::t -> h!="#" && aux t in if aux w then 1 else 0
        | ["#"], [n] -> if tmp+1 = n then 1 else 0
        | ["#"], _ -> 0
        | "#"::tw, n::_ -> if tmp+1 > n then 0 else doPart1 tw num (tmp+1) false
        | "."::tw, n::tnum -> if first then doPart1 tw num 0 first else if tmp < n then 0 else doPart1 tw tnum 0 true
        | "?"::tw, _ -> let r1 = doPart1 ("#"::tw) num tmp false and r2 = doPart1 ("."::tw) num tmp first in r1+r2 
        | _ -> 0


(*?#??????##? 1,1,2*)
;;  

let doPart2 w num =
        let nw = List.length w in let nnum = List.length num in 
        let opt = Array.make (nw*10) (Array.make (nnum*10) (Array.make (nw*nnum) (-1))) in
        let rec aux w num i1 i2 tmp = let value = opt.(i1).(i2).(tmp) in if value != -1 then value else 
        match w, num with 
                | [], [] -> if tmp = 0 then 1 else 0
                | [], [n] -> if tmp = n then 1 else 0
                | [], _ -> 0
                | "?"::tw, _ -> let r1 = aux tw num (i1+1) i2 (tmp+1) in 
                        if tmp = 0 then let res = (aux tw num (i1+1) i2 0) + r1 in  (opt.(i1).(i2).(tmp) <- res; res)
                        else
                        (match num with 
                                | n::tnum -> if n = tmp then let res = (aux tw tnum (i1+1) (i2+1) 0) + r1 in  (opt.(i1).(i2).(tmp) <- res; res) else (opt.(i1).(i2).(tmp) <- r1; r1)
                                | [] ->  (opt.(i1).(i2).(tmp) <- r1; r1)
                        )
                | "."::tw, _::_ -> if tmp = 0 then let res = aux tw num (i1+1) i2 (tmp+1) in (opt.(i1).(i2).(tmp) <- res; res)                           else 
                        (match num with 
                                | n::tnum -> if n = tmp then let res = aux tw tnum (i1+1) (i2+1) 0 in 
                                (opt.(i1).(i2).(tmp) <- res; res)
                                else (opt.(i1).(i2).(tmp) <- 0; 0) 
                                | _ -> 0)
                | "#"::tw, _::_ -> let res = aux tw num (i1+1) i2 (tmp+1) in (opt.(i1).(i2).(tmp) <- res; res)
                | _, _ -> 0
        
        in aux w num 0 0 0

;;


let pars file_name =
  try
    let channel = open_in file_name in
    let rec read_lines res1 res2 index =
      try
        let line = input_line channel in
        let w, num = match split line " " with | h1::h2::[] -> (split h1 ""), (List.map int_of_string (split h2 ",")) | _ -> [], [] in 
        let p1 = doPart1 w num 0 true in 
        let w2 = w@["?"]@w@["?"]@w@["?"]@w@["?"]@w in 
        let num2 = num@num@num@num@num in
        let p2 = doPart2 w2 num2 in (print_int p1; print_endline "";      
        read_lines (res1+p1) (res2+p2) (index+1))
      with
      | End_of_file -> (close_in channel; res1, res2)
    in
    read_lines 0 0 1
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;
                      
let a, b = (pars "inputs/input_test12.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input12.txt") in (print_int a; print_string " "; print_int b; print_endline "");;
