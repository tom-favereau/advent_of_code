open AOC_ocaml_lib.Utils

module Day10 : sig 
  
  val part1 : string -> int

  val part2 : string -> int

end = struct 
  type from = Up | Down | Right | Left 
  
  let dfs (is, js) mat = 
    let rec aux (i, j) fr tmp = 
      match mat.(i).(j), fr with
      | "|", Up -> aux (i+1, j) Up ((i, j)::tmp)
      | "|", Down -> aux (i-1, j) Down ((i, j)::tmp)
      | "-", Right -> aux (i, j-1) Right ((i, j)::tmp)
      | "-", Left -> aux (i, j+1) Left ((i, j)::tmp)
      | "7", Left ->aux (i+1, j) Up ((i, j)::tmp)
      | "7", Down ->aux (i, j-1) Right ((i, j)::tmp)
      | "J", Up ->aux (i, j-1) Right ((i, j)::tmp) 
      | "J", Left ->aux (i-1, j) Down ((i, j)::tmp) 
      | "L", Right -> aux (i-1, j) Down ((i, j)::tmp) 
      | "L", Up -> aux (i, j+1) Left ((i, j)::tmp) 
      | "F", Right -> aux (i+1, j) Up ((i, j)::tmp)
      | "F", Down -> aux (i, j+1) Left ((i, j)::tmp)
      | "S", _ -> tmp 
      | _ -> []
    in 
    let r1 = aux (is+1,js) Up [] in      
    let r2 = aux (is-1,js) Down [] in      
    let r3 = aux (is,js-1) Right [] in      
    let r4 = aux (is,js+1) Left [] in      
    let r = List.filter (fun e -> List.length e > 0) [r1; r2; r3; r4] in
    List.hd r
    

  let aire lp = 
    let xf, yf = List.hd lp in
    let rec aux lp tmp = match lp with
      | [] -> tmp
      | [a] -> let x, y = a in tmp+(x*yf-xf*y)
      | a::b::tail -> let x1, y1 = a and x2, y2 = b in  
                      aux (b::tail) (tmp+(x1*y2-x2*y1))
    in let a = aux lp 0 in if a < 0 then -a/2 else a/2
  


  let part1 file_name = 
    let matrix = match Pars.read_to_matrix file_name with 
      | None -> failwith "parsing"
      | Some mat -> mat
    in
    let ps = ref (0, 0) in
    Array.iteri (fun i line -> 
      Array.iteri (fun j e -> 
        if e = "S" then ps := (i, j)) line) matrix; 
    let ch = dfs !ps matrix in 
    (List.length ch)/2+1

  let part2 file_name = 
    let matrix = match Pars.read_to_matrix file_name with 
      | None -> failwith "parsing"
      | Some mat -> mat
    in 
    let ps = ref (0, 0) in 
    Array.iteri (fun i line -> 
      Array.iteri (fun j e -> 
        if e = "S" then ps := (i, j)) line) matrix; 
    let ch = dfs !ps matrix in
    let r = aire ch in
    r-(List.length ch)/2 + 1 
  
end


let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else begin
    let file_name = Sys.argv.(1) in
    let r1 = Day10.part1 file_name in 
    let r2 = Day10.part2 file_name in 
    (print_int r1; print_newline (); print_int r2; print_newline ())
  
  end

