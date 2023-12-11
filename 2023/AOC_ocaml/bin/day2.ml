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

let part1 file_name =
  try
    let channel = open_in file_name in
    let rec read_lines res index  =
      try
        let line = input_line channel in
        let liste = match split line ":" with | _::head::_ -> split head ";" | _ -> [] in
        if doPart1 liste then read_lines (res+index) (index+1) else read_lines res (index+1)
      with
      | End_of_file -> (close_in channel; res)
    in
    read_lines 0 1
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0)
;;


                      

print_int (part1 "inputs/input_test02.txt");;
print_endline "";;
print_int (part1 "inputs/input02.txt");;
print_endline "";;
