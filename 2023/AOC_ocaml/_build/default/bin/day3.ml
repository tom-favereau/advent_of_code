
let split s exp = Str.split (Str.regexp exp) s;;


let ll_to_mat ll = 
  let arr = Array.make_matrix (List.length ll) (List.length (List.hd ll)) "" in 
  let rec loop2 i j = function
    | [] -> ()
    | h::t -> (arr.(i).(j) <- h; loop2 i (j+1) t)
  in 
  let rec loop1 i = function 
    | [] -> ()
    | h::t -> (loop2 i 0 h; loop1 (i+1) t) 
  in (loop1 0 ll; arr)
;;


let isSymbol s = 
  let numbers = [|"0"; "1"; "2"; "3"; "4"; "5"; "6"; "7"; "8"; "9"; "."|] in
  let res = ref true in 
  for i = 0 to (Array.length numbers) - 1 do
    if numbers.(i) = s then res := false 
  done;
  !res
;;
(*
let print_mat m = for i = 0 to (Array.length m)-1 do for j = 0 to (Array.length m.(0)-1) do
  print_string m.(i).(j) done; print_newline () done; print_newline ();;  
*)
(*
let doPart1 arr = 
  let n = Array.length arr in
  let m = Array.length arr.(0) in
  let res = ref 0 in
  for i = 0 to n-1 do 
    for j = 0 to m-1 do 
      if isSymbol arr.(i).(j) then 
        let subArr = Array.make_matrix 3 7 "" in
        for k = 0 to 2 do for l = 0 to 6 do
          try subArr.(k).(l) <- arr.(i+k-1).(j+l-3) 
          with | _ -> subArr.(k).(l) <- "."
        done; done; 
        let aux sa = let num = ref 0 in 
          for k = 0 to 2 do for l = 0 to 6 do
            if not (isSymbol sa.(k).(l)) && sa.(k).(l) <> "." then 
              num := !num*10 + (int_of_string sa.(k).(l))  
            else if not (isSymbol sa.(k).(l)) then ((*if !num <> 0 then (print_int !num; print_string " ");*) res := !res + !num; num := 0) 
          done; done;
        in (print_mat subArr; aux subArr)
    done;
  done;
  !res 
                
;;
*)

let doPart1 arr = 
  let n = Array.length arr and m = Array.length arr.(0) in
  let res = ref 0 in
  let num = ref 0 in
  let isValid = ref false in
  let valid k l = let res = ref false in for i = -1 to 1 do for j = -1 to 1 do 
    try if isSymbol arr.(k+i).(l+j) then res := true with _ -> () done; done; !res in  
  for i = 0 to n-1 do for j = 0 to m-1 do
    try (let tmp = int_of_string arr.(i).(j) in 
        num := !num*10 + tmp; if valid i j then isValid := true)
    with _  -> (if !isValid then res := !res + !num; num := 0; isValid := false) 
  done; (if !isValid then res := !res + !num; isValid := false; num := 0) done; !res
;;



let auxPart2 arr = 
  let n = Array.length arr and m = Array.length arr.(0) in
  let res = ref 1 in
  let count = ref 0 in
  let num = ref 0 in
  let isValid = ref false in
  let valid k l = let res = ref false in for i = -1 to 1 do for j = -1 to 1 do 
    try if isSymbol arr.(k+i).(l+j) then res := true with _ -> () done; done; !res in  
  for i = 0 to n-1 do for j = 0 to m-1 do
    try (let tmp = int_of_string arr.(i).(j) in 
        num := !num*10 + tmp; if valid i j then isValid := true)
    with _  -> (if !isValid then (res := !res * !num; count := !count+1); num := 0; isValid := false) 
  done; (if !isValid then (res := !res * !num; count := !count + 1);isValid := false; num := 0) done; 
  (!res, !count)
;;


let doPart2 arr = 
  let n = Array.length arr and m = Array.length arr.(0) in 
  let res = ref 0 in
  for i = 0 to n-1 do for j = 0 to m-1 do 
    if arr.(i).(j) = "*" then 
      let subArr = Array.make_matrix 3 7 "." in 
      for k = -1 to 1 do for l = -3 to 3 do 
        if not (isSymbol arr.(k+i).(l+j)) then subArr.(k+1).(l+3) <- arr.(k+i).(l+j) 
      done; done; (subArr.(1).(3) <- "*");    
      let tmp, count = auxPart2 subArr in (if count = 2 then res := !res + tmp;)
  done; done; !res
;;


let pars file_name =
  try
    let channel = open_in file_name in
    let rec read_lines mat  =
      try
        let line = input_line channel in
        let arr = split line "" in
        read_lines (arr::mat) (*c'est super moche mais je comprend par pourquoi le compilateur accepte pas arr::mat*)  
      with
      | End_of_file -> (close_in channel; mat)
    in
    let mat = ll_to_mat (read_lines []) in
    let r1 = doPart1 mat in 
    let r2 = doPart2 mat in
    (r1, r2)

  with
  | Sys_error err ->( print_endline ("Error: " ^ err); 0, 0)
;;
                      
let a, b = (pars "inputs/input_test03.txt") in (print_int a; print_string " "; print_int b;  print_endline "");;
let a, b = (pars "inputs/input03.txt") in (print_int a; print_string " "; print_int b; print_endline "");;

