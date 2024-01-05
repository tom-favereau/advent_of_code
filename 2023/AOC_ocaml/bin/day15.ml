open AOC_ocaml_lib.Utils

module Day15 : sig 
  val part1 : string -> int

  val part2 : string -> int
end = struct 

  let hash l =
    let rec aux ll tmp res = match ll with
        | [] -> res+tmp
        | ","::t -> aux t 0 (res+tmp)
        | a::t -> let v = Char.code a.[0] in let ntmp = (17*(tmp+v)) mod 256 in aux t ntmp res
    in aux l 0 0

  let part1 file_name = 
    let line = match Pars.read_lines file_name with 
      | None -> failwith "parsing"
      | Some lines -> List.hd lines
    in
    let data = Pars.split line "," in
    List.fold_left (fun acc e -> acc + hash (Pars.split e "") 
    ) 0 data

  let part2 file_name = 
    let line = match Pars.read_lines file_name with 
      | None -> failwith "parsing"
      | Some lines -> List.hd lines 
    in 
    let data = Pars.split line "," in
    let table = Array.make 256 [] in
    List.iter (fun e -> 
      let size = String.length e in
      if e.[size-1] = '-' then 
        let e' = Str.string_before e (size-1) in 
        let hash_e' = Pars.split e' "" |> hash in
        let rec del = function
          | [] -> []
          | (e, value)::t -> if e = e' then t else (e, value)::(del t)
        in table.(hash_e') <- del table.(hash_e');
      else 
        let sep = Pars.split e "=" in 
        let e' = (sep |> List.hd) in
        let h = Pars.split e' "" |> hash in
        let value = sep |> List.tl |> List.hd |> int_of_string in
        let rec add = function 
          | [] -> [(e', value)]
          | (e, v)::t -> if e = e' then (e, value)::t else (e, v)::(add t)  
        in table.(h) <- add table.(h);
    ) data; 
    let res, _ = Array.fold_left (fun (res, ind) e -> 
      let acc, _ = List.fold_left (fun (acc, ind') (_, value)   ->
        (acc + value * ind' * ind, ind'+1)
      ) (0, 1) e in
      (res + acc, ind+1)
    ) (0, 1) table in 
    res 
  
end

  
let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else begin
    let file_name = Sys.argv.(1) in
    let r1 = Day15.part1 file_name in 
    let r2 = Day15.part2 file_name in 
    (print_int r1; print_newline (); print_int r2; print_newline ())
  
  end

