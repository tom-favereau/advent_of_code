open AOC_ocaml_lib.Utils 

module Day14 : sig 
  val part1 : string -> int

  val part2 : string -> int
end = struct 

  let rotate : string array array -> string array array = fun matrix ->
    let n = Array.length matrix and m = Array.length matrix.(0) in
    let res = Array.make_matrix m n "" in
    for i = 0 to n-1 do for j = 0 to m-1 do
      res.(j).(n-1-i) <- matrix.(i).(j)
    done; done; 
    res

  let roll : string array array -> string array array = fun matrix ->
    let n = Array.length matrix and m = Array.length matrix.(0) in
    let current = Array.make n 0 in
    let res = Array.copy matrix in
    for i = 0 to n-1 do for j = 0 to m-1 do 
      let elem = matrix.(i).(j) in
      if elem = "#" then current.(j) <- i+1
      else if elem = "O" then (
        res.(i).(j) <- ".";
        res.(current.(j)).(j) <- "O";
        current.(j) <- current.(j)+1)
    done; done;
    res

  let score : string array array -> int = fun matrix -> 
    let n = Array.length matrix and m = Array.length matrix.(0) in
    let res = ref 0 in
    for i = 0 to n-1 do for j = 0 to m-1 do 
      let elem = matrix.(i).(j) in 
      if elem = "O" then res := !res + (n-i-1)+1
    done; done;
    !res
  
  let part1 file_name = 
    let matrix = match Pars.read_to_matrix file_name with 
      | None -> failwith "parsing" 
      | Some mat -> mat 
    in
    matrix |> roll |> score 

  let part2 file_name = 
    let arr_to_string arr = Array.fold_left (fun acc elem -> acc ^ elem) "" arr in
    let mat_to_string mat = Array.fold_left (fun acc elem -> acc ^ (arr_to_string elem)) "" mat in
    let matrix = match Pars.read_to_matrix file_name with 
      | None -> failwith "parsing"
      | Some mat -> ref mat
    in
    let config = Hashtbl.create 1000 in 
    let scores = Hashtbl.create 1000 in
    let res = ref 0 in
    let i = ref (-1) in
    let continuer = ref true in
    while !i < 1000000000  && !continuer do 
      for _ = 0 to 3 do begin matrix := roll !matrix; matrix := rotate !matrix end done;
      i := !i+1;
      let s = mat_to_string !matrix in
      if Hashtbl.mem config s then
        let index = Hashtbl.find config s in
        let tmp = (1000000000- !i) mod (!i-index)-1 in
        (
        res := Hashtbl.find scores (tmp+index);
        continuer := false)
      else (Hashtbl.add config s !i; Hashtbl.add scores !i (score !matrix))
    done;
    !res


end



  
let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else begin
    let file_name = Sys.argv.(1) in
    let r1 = Day14.part1 file_name in 
    let r2 = Day14.part2 file_name in 
    (print_int r1; print_newline (); print_int r2; print_newline ())
  
  end

