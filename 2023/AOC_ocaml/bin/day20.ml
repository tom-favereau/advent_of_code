open AOC_ocaml_lib.Utils

module Day20 : sig 
  val part1 : string -> int 

end = struct
(*
  type info = {
    mutable on : int;
    mutable prev : bool list;
  }
  
  let count graph types = 
  *)  

  let part1 file_name = 
    let lines = match Pars.read_lines file_name with 
      | None -> failwith "parsing" 
      | Some lines -> lines 
    in
    let graph = Hashtbl.create (List.length lines) in    
    let types = Hashtbl.create (List.length lines) in
    List.iter (fun line -> 
      let sep = Pars.split line " -> " in
      let node = if List.hd sep = "broadcaster" then List.hd sep 
      else Str.string_after (List.hd sep) 1 in
      let dest = Pars.split (sep |> List.tl |> List.hd) ", " in
      List.iter (fun d -> Hashtbl.add graph node d) dest;
      Hashtbl.add types node (List.hd sep).[0];
    ) lines;
    Hashtbl.iter (fun key value -> Printf.printf "%s -> %s \n" key value) graph;
    0

end


let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else begin
    let file_name = Sys.argv.(1) in
    let r1 = Day20.part1 file_name in 
    (*let r2 = Day7.part2 file_name in 
    *)
    (print_int r1; print_newline ();)
  end

