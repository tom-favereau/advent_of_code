
let split s exp = Str.split (Str.regexp exp) s;;


let list_to_arr l = 
  let res = Array.make_matrix (List.length l) (List.length (List.hd l)) "" in
  let rec aux i = function 
    | [] -> ()
    | h::t -> let rec loop j = function | [] -> () | a::q -> (res.(i).(j) <- a; loop (j+1) q) 
      in loop (i+1) h  
  in (aux 0 l; res)
;;
    


let pars file_name =
  try
    let channel = open_in file_name in
    let rec read_lines liste_lines  =
      try
        let line = input_line channel in
        let list = split line "" in
        read_lines list::liste_lines
      with
      | End_of_file -> (close_in channel; liste_lines)
    in
    let liste_line = read_lines [] in 
    let arr = list_to_arr liste_line in
    for i = 0 to Array.length arr do for j = 0 to Array.length arr do 
        print_string arr.(i).(j) done; done;
    0, 0

  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;
                      
let a, b = (pars "inputs/input_test18.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input18.txt") in (print_int a; print_string " "; print_int b; print_endline "");;
