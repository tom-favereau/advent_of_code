open AOC_ocaml_lib.Utils



module Day20 : sig 
  val part1 : string -> int 

end = struct

  type info = 
    | Flip of bool 
    | Conj of (string * bool) list

  module MMap = Map.Make(struct 
    type t = string
    let compare = compare
  end)
  
  type signal = Hight of string | Low of string 

  let count graph types = 
    let queue = Queue.create () in 
    let low = ref 0 and hight = ref 0 in 
    Queue.add (Hashtbl.find graph "broadcaster" |> List.map (fun e -> Low e)) queue; 
    let actualise (elem:string) = Hashtbl.find graph elem |> 
      List.iter (fun e -> match MMap.find e types with 
        | Conj ll -> let nl = List.map (function | 
            e -> if e = elem then e) ll in () 
        | _ -> () 
      ) in 
    let rec aux types = match Queue.take_opt queue with 
      | None -> ()
      | Some elems -> List.iter (function 
        | Hight e ->
          begin match MMap.find e types with
            | Conj ll -> ()
            | _ -> ()
          end
        | Low e -> 
          begin match MMap.find e types with 
            | Flip b ->
              Hashtbl.find graph e |> List.iter (fun e' -> );
              if b then (low := !low + 1; types |> MMap.add e (Flip false) |> aux)
              else (hight := !hight + 1; types |> MMap.add e (Flip false) |> aux)
            | Conj ll when List.fold_left (fun acc e -> e && acc) true ll -> ()
            | Conj _ -> ()
          end
      ) elems
    in
    aux types;
    !low * !hight

  let part1 file_name = 
    let lines = match Pars.read_lines file_name with 
      | None -> failwith "parsing" 
      | Some lines -> lines 
    in
    let graph = Hashtbl.create (List.length lines) in    
    let types = MMap.empty |>
    List.fold_right (fun line typ -> 
      let sep = Pars.split line " -> " in
      let node = if List.hd sep = "broadcaster" then List.hd sep 
      else Str.string_after (List.hd sep) 1 in
      let dest = Pars.split (sep |> List.tl |> List.hd) ", " in
      Hashtbl.add graph node dest;
      if (List.hd sep).[0] = '%' then MMap.add node (Flip false) typ 
      else MMap.add node (Conj []) typ
    ) lines in
    Hashtbl.iter (fun key value -> Printf.printf "%s -> " key; List.iter (fun e -> Printf.printf "%s, " e) value; print_newline() ) graph;
    let types = Hashtbl.fold (fun key value acc -> match MMap.find key acc with 
      | Flip _ -> acc
      | Conj _ -> MMap.add key (Conj (List.map (fun e -> (e, false)) value)) acc 
    ) graph types in  
    count graph types

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

