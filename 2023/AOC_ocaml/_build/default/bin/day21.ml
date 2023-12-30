open AOC_ocaml_lib.Utils 


module Day21 : sig 
  val part1 : string -> int -> int

  val part2 : string -> int

end = struct 

  let bfs : string array array -> int -> (int * int) =
    fun matrix max_step -> 
    let queue = Queue.create () in 
    let size = Array.length matrix in 
    let ps = size / 2 in
    let module Mset = Set.Make(struct  
      type t = (int*int*int) 
      let compare = compare 
      end) 
    in
    let vis = Hashtbl.create max_step in
    Queue.add (ps, ps, 0) queue;
    let rec bfs_aux set = match Queue.take_opt queue with 
      | None -> set
      | Some (i, j, step) ->  
        let i' = if (i mod size) < 0 then (i mod size) + size else i mod size in
        let j' = if (j mod size) < 0 then (j mod size) + size else j mod size in 
        let elem = matrix.(i' mod size).(j' mod size) in
        if elem = "#" then bfs_aux set
        else if step >= max_step then bfs_aux (Mset.add (i, j, step) set) 
        else if not (Hashtbl.mem vis (i, j)) then (Queue.add_seq queue 
        (List.to_seq [(i+1, j, step+1); (i-1, j, step+1); (i, j+1, step+1); (i, j-1, step+1)]);
        Hashtbl.add vis (i, j) true;
        bfs_aux (Mset.add (i, j, step) set))    
        else bfs_aux set
    in 
    let set = bfs_aux Mset.empty in
    let odd = ref 0 and eve = ref 0 in 
    let mem = Hashtbl.create  0 in
    Mset.iter (fun (a, b, i) -> if not (Hashtbl.mem mem (a, b)) then 
    if i mod 2 = 0 then eve := !eve+1 else odd := !odd+1; Hashtbl.add mem (a, b) true) set; 
    (!eve, !odd)

  let part1 : string -> int -> int = fun file_name step -> 
    let lines = match Pars.read_lines file_name with 
    | Some lines -> lines
    | None -> failwith "parsing"
    in
    let size = List.length lines in
    let matrix = Array.make_matrix size size "" in 
    List.iteri (fun i s -> 
      List.iteri (fun j e -> matrix.(i).(j) <- e) (Pars.split s "")
      ) lines; 
    let res, _ = bfs matrix step in
    res
    

  let part2 : string -> int = fun file_name -> 
    let lines = match Pars.read_lines file_name with 
    | Some lines -> lines
    | None -> failwith "parsing"
    in
    let size = List.length lines in
    let matrix = Array.make_matrix size size "" in 
    List.iteri (fun i s -> 
      List.iteri (fun j e -> matrix.(i).(j) <- e) (Pars.split s "")
      ) lines;
    let size = Array.length matrix in
    let x = 26501365/size in
    let rest = 26501365 mod size in
    let x1 = 0 and x2 = 2 and x3 = 4 in
    let foi = float_of_int in
    let _, p1 = bfs matrix (x1*size + rest) in
    let _, p2 = bfs matrix (x2*size+rest) in
    let _, p3 = bfs matrix (x3*size+rest) in
    let y1 = foi p1 and y2 = foi p2 and y3 = foi p3 in
    let x1 = foi x1 and x2 = foi x2 and x3 = foi x3 in
    let x = foi x in
    int_of_float (
    ((x -. x2)*.(x-.x3)*.y1/.((x1-.x2)*.(x1-.x3))) +.
    ((x -. x1)*.(x-.x3)*.y2/.((x2-.x1)*.(x2-.x3))) +.
    ((x -. x1)*.(x-.x2)*.y3/.((x3-.x1)*.(x3-.x2)))
    )

end


  
let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else begin
    let file_name = Sys.argv.(1) in
    let step = if file_name = "inputs/input21.txt" then 64 else 6 in
    let r1 = Day21.part1 file_name step in 
    let r2 = Day21.part2 file_name in 
    (print_int r1; print_newline (); print_int r2; print_newline ())
  end

