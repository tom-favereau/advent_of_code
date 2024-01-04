open AOC_ocaml_lib.Utils

module Day9 : sig 

  val part1 : string -> int

  val part2 : string -> int
  
end = struct 

  let rec predict1 line = 
    let rec aux tmp zero = function 
      | [] -> tmp, zero 
      | [_] ->  tmp, zero
      | h1::h2::t -> aux ((h2-h1)::tmp) (zero && (h2-h1)=0) (h2::t)
    in 
    let l, zero = aux [] true line in
    if zero then line |> List.rev |> List.hd 
    else (line |> List.rev |> List.hd) + predict1 (List.rev l)

  let rec predict2 line =   
    let rec aux tmp zero = function 
      | [] -> tmp, zero 
      | [_] ->  tmp, zero
      | h1::h2::t -> aux ((h2-h1)::tmp) (zero && (h2-h1)=0) (h2::t)
    in 
    let l, zero = aux [] true line in
    if zero then line |> List.hd 
    else (line |> List.hd) - predict2 (List.rev l) 


  let part1 file_name = 
    let lines = 
      match Pars.read_lines file_name with 
      | None -> failwith "parsing" 
      | Some lines -> lines 
    in
    List.fold_left (fun acc line -> 
      acc+(predict1 (List.map int_of_string (Pars.split line " ")))) 0 lines

  let part2 file_name = 
    let lines = 
      match Pars.read_lines file_name with 
      | None -> failwith "parsing" 
      | Some lines -> lines 
    in
    List.fold_left (fun acc line -> 
      acc+(predict2 (List.map int_of_string (Pars.split line " ")))) 0 lines


end 



  
let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else begin
    let file_name = Sys.argv.(1) in
    let r1 = Day9.part1 file_name in 
    let r2 = Day9.part2 file_name in 
    (print_int r1; print_newline (); print_int r2; print_newline ())
  
  end

