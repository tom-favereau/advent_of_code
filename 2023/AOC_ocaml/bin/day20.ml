open AOC_ocaml_lib.Utils

module Dgraph = Graph.Imperative.Digraph.ConcreteBidirectional (struct
  type t = string
  let compare = compare
  let hash = Hashtbl.hash
  let equal = (=)   
end)


module Day20 : sig 
  val part1 : string -> int 
end = struct
  (*
  type state = {name : string; sym : string; up : bool}
  *)
  let print_graph graph =
    Dgraph.iter_vertex (fun vertex -> Printf.printf "Sommet : %s\n" vertex) graph;
    Dgraph.iter_edges (fun source target -> Printf.printf "Arête : %s -> %s\n" source target) graph

  (*
  let bfs : Dgraph.t -> (int * int) = fun graph -> 
    let queue = Queue.create () in  
    Queue.add "broadcaster" queue;
    (0, 0)
   *)
  let part1 : string -> int = fun file_name -> 
    let lines = match Pars.read_lines file_name with 
          | None -> failwith "name file"
          | Some thing -> thing
    in
    let graph = Dgraph.create ~size:(List.length lines) () in
    let vert = Hashtbl.create (List.length lines) in
    List.iter (fun elem ->
    let node, dest = match (Pars.split elem " -> ") with 
      | [nod; des] ->  let label = if nod <> "broadcaster" then Str.string_before nod 1 else "" in 
                        let nod' = if nod <> "broadcaster" then Str.string_after nod 1 else nod in 
        (Hashtbl.add vert nod' label ; Dgraph.add_vertex graph nod'; nod', des) 
      | l -> (List.iter print_string l; failwith "parsing") 
    in  
    List.iter (fun elem' -> 
        (if not (Dgraph.mem_vertex graph elem') 
          then (Dgraph.add_vertex graph elem'); 
        Dgraph.add_edge graph node elem') 
      ) (Pars.split dest ", ")
    ) lines;
    (print_graph graph;
    let module Bfs = Graph.Traverse.Bfs(Dgraph) in
    0)

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

