(* day1.ml *)

module Day1 = struct 

let string_to_list str = str |> String.to_seq |> List.of_seq;;

let rec traitement line = match line with 
        | '0'::_ -> 0
        | '1'::_ -> 1
        | '2'::_ -> 2
        | '3'::_ -> 3
        | '4'::_ -> 4
        | '5'::_ -> 5
        | '6'::_ -> 6
        | '7'::_ -> 7
        | '8'::_ -> 8
        | '9'::_ -> 9
        | _::q -> traitement q 
        | _ -> 0

;;


let part1 file_name =
  try
    let channel = open_in file_name in
    let rec read_lines res  =
      try
        let line = input_line channel in
        let li = string_to_list line in
        let num1 = traitement li in let num2 = traitement (List.rev li) in
        let tmp = num1*10 + num2 in
        read_lines res+tmp 
      with
      | End_of_file -> (close_in channel; res)
    in
    read_lines 0 
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0)
;;

print_int (part1 "inputs/input_test01.txt");;
print_endline "";;
print_int (part1 "inputs/input01.txt");;
print_endline "";;


end
