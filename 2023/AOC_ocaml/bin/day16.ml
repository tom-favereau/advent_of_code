let split s exp = Str.split (Str.regexp exp) s;;



let pars file_name =
  try
    let channel = open_in file_name in
    let rec read_lines tmp  =
      try
        let line = input_line channel in
        let arr = Array.make (String.lenght line) '' in
        for i = 0 to String.lenght line-1 do
            arr.(i) <- line.[i]
        done;
        read_lines arr::tmp
      with
      | End_of_file -> (close_in channel; tmp)
    in
    let tmp = read_lines []
    let arr = Array.make (List.lenght tmp) (Array.make (Array.length (List.head tmp)) '') in
    let toArr c = function | [] -> () | h::t -> (arr.(c) <- h; toArr (tmp+1) t) in
    toArr tmp
  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;

type info = {i : int; j : int; dir : string};;

let propagation arr =
    let ener = Array.make n (Array.make m false) in
    let visited = Hashtbl.create 10 in
    let rec dfs i j dir =
        let k = {i:i; j:j; dir;dir} in
        if Hashtbl.mem visited k then ()
        else 

let p1 file_name =
    let arr = pars file_name in
    let n = Array.length arr in
    let m = Array.length arr.(0) in
    let dfs i j dir =


let a, b = (pars "inputs/input_test02.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input02.txt") in (print_int a; print_string " "; print_int b; print_endline "");;
