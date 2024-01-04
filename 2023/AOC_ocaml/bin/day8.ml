open AOC_ocaml_lib.Utils

module Day8 : sig 
  val part1 : string -> int
end = struct 

  let parcours graph ins =
    let n = Array.length ins in
    let rec aux tmp index = 
      let choice = Hashtbl.find graph tmp in  
      let dir = ins.(index mod n) in 
      if choice.(dir) = "ZZZ" then index 
      else aux choice.(dir) (index+1)
    in
    aux "AAA" 0 + 1
  
  let part1 file_name =
    let contents = 
    match Pars.read file_name with 
    | None -> failwith "parsing"
    | Some contents -> contents 
    in 
    let sep = Pars.split contents "\n\n" in
    let ins = Pars.split (List.hd sep) "" |> Array.of_list |> 
    Array.map (fun e -> if e = "L" then 0 else 1) in 
    let graph = sep |> List.length |> Hashtbl.create in
    List.iter (fun line -> 
        let tmp = Pars.split line " = (" in
        let node = List.hd tmp in 
        let dest = Pars.split (List.hd (List.tl tmp)) ", " in
        let l = (List.hd dest) in
        let r = (Str.string_before (List.hd (List.tl dest)) 3) in
        Hashtbl.add graph node [|l; r|];
    ) (Pars.split (sep |> List.tl |> List.hd) "\n");
    parcours graph ins 
  
end 

  
let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else begin
    let file_name = Sys.argv.(1) in
    let r1 = Day8.part1 file_name in 
    (*let r2 = Day9.part2 file_name in 
    *)(print_int r1; print_newline ())(*; print_int r2; print_newline ())
  *)
  end

