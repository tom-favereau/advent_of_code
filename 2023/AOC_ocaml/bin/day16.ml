open AOC_ocaml_lib.Utils 

module Day16 : sig 
  val part1 : string -> int
  val part2 : string -> int
end = struct 
  type dir = Up | Down | Right | Left
  
  let dfs ps ds matrix = 
    let n = Array.length matrix in
    let visited = Hashtbl.create (n*n) in
    let count = Hashtbl.create (n*n) in
    let rec aux (i, j) d = 
      if not ((Hashtbl.mem visited (i, j, d)) || (i < 0) || (i >= n) || (j < 0) || (j >= n)) then 
      begin   
      if not (Hashtbl.mem count (i, j)) then Hashtbl.add count (i, j) 0;  
        Hashtbl.add visited (i, j, d) 0;
        match matrix.(i).(j), d with 
        | ("." | "|"), Up -> aux (i-1, j) Up 
        | ("." | "|"), Down -> aux (i+1, j) Down
        | ("." | "-"), Right -> aux (i, j+1) Right
        | ("." | "-"), Left -> aux (i, j-1) Left
        | "|", (Right | Left) -> (aux (i-1, j) Up; aux (i+1, j) Down)
        | "-", (Up | Down) -> (aux (i, j+1) Right; aux (i, j-1) Left)
        | ("/", Up | "\\", Down) -> aux (i, j+1) Right
        | "/", Down | "\\", Up -> aux (i, j-1) Left
        | "/", Left | "\\", Right -> aux (i+1, j) Down
        | "/", Right | "\\", Left -> aux (i-1, j) Up
        | _ -> failwith "patern matching"
    end
    in (aux ps ds; 
    Hashtbl.length count)   
    
  
  let part1 file_name = 
    let matrix = match Pars.read_to_matrix file_name with 
      | None -> failwith "parsing"
      | Some matrix -> matrix
    in
    dfs (0, 0) Right matrix 

  let part2 file_name = 
    let matrix = match Pars.read_to_matrix file_name with 
      | None -> failwith "parsing"
      | Some matrix -> matrix 
    in 
    let n = Array.length matrix in
    [
    List.init n (fun i -> dfs (0, i) Down matrix) |> List.fold_left max 0;
    List.init n (fun i -> dfs (n-1, i) Up matrix) |> List.fold_left max 0;
    List.init n (fun i -> dfs (i, 0) Right matrix) |> List.fold_left max 0;
    List.init n (fun i -> dfs (i, n-1) Left matrix) |> List.fold_left max 0;
    ] |> List.fold_left max 0

end


let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else 
    let file_name = Sys.argv.(1) in
    let r1 = Day16.part1 file_name in 
    let r2 = Day16.part2 file_name in 
    (print_int r1; print_newline (); print_int r2; print_newline ())
  

