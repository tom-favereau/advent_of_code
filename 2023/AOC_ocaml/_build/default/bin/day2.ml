(*open Str;;*)

(*let string_to_list str = str |> String.to_seq |> List.of_seq;;*)

let split s exp = Str.split (Str.regexp exp) s;;


let last str =
  let length = String.length str in
  if length > 0 then
    Some str.[length - 1]
  else
    None
;;

let rec doPart1 l = match l with 
        | [] -> true 
        | h::t ->
            let rec aux ll r g b = if r > 12 || g > 13 || b > 14 then false else match ll with  
                | [] -> true
                | a::q -> let num = match split a " " with | n::_ -> int_of_string n | _ -> 0 in
                match last a with 
                        | Some 'e' -> aux q r g (b+num) 
                        | Some 'd' -> aux q (r+num) g b
                        | Some 'n' -> aux q r (g+num) b
                        | _ -> false
            in (aux (split h ",") 0 0 0) && (doPart1 t)
;;

let rec doPart2 l = match l with 
        | [] -> 0, 0, 0
        | h::t ->
            let rec aux ll r g b = match ll with  
                | [] -> r, g, b
                | a::q -> let num = match split a " " with | n::_ -> int_of_string n | _ -> 0 in
                match last a with 
                        | Some 'e' -> aux q r g (b+num) 
                        | Some 'd' -> aux q (r+num) g b
                        | Some 'n' -> aux q r (g+num) b
                        | _ -> 0, 0, 0
            in 
            let r1, g1, b1 = (aux (split h ",") 0 0 0) in 
            let r2, g2, b2 = (doPart2 t) in max r1 r2, max g1 g2, max b1 b2
;;

let pars file_name =
  try
    let channel = open_in file_name in
    let rec read_lines res1 res2 index  =
      try
        let line = input_line channel in
        let liste = match split line ":" with | _::head::_ -> split head ";" | _ -> [] in
        let p1 = doPart1 liste in 
        let p2 = match doPart2 liste with r, g, b -> r*b*g in
        if p1 then read_lines (res1+index) (res2+p2) (index+1) else read_lines res1 (res2+p2) (index+1)
      with
      | End_of_file -> (close_in channel; res1, res2)
    in
    read_lines 0 0 1
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;
                      
let a, b = (pars "inputs/input_test02.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input02.txt") in (print_int a; print_string " "; print_int b; print_endline "");;
