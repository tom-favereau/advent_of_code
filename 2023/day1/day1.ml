let string_to_list str = str |> String.to_seq |> List.of_seq;;

let rec traitement line = match line with 
        | '0'::q -> 0
        | '1'::q -> 1
        | '2'::q -> 2
        | '3'::q -> 3
        | '4'::q -> 4
        | '5'::q -> 5
        | '6'::q -> 6
        | '7'::q -> 7
        | '8'::q -> 8
        | '9'::q -> 8
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
        let tmp = (traitement li)*10 + (traitement (List.rev li)) in  
        read_lines res+tmp 
      with
      | End_of_file -> (close_in channel; res)
    in
    read_lines 0 
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0)
;;




                      

print_int (part1 "input2.txt");;
