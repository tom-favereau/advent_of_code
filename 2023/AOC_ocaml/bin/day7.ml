open AOC_ocaml_lib.Utils


module Day7 : sig 

  val part1 : string -> int 

  val part2 : string -> int

end = struct 

  type hand = {
    cards : string list;
    points : int; 
  }

  let cmp_cards (h1 : hand) (h2 : hand) : int = 
    let card_value = Hashtbl.create 10 in 
    Hashtbl.add card_value "A" 13;
    Hashtbl.add card_value "K" 12;
    Hashtbl.add card_value "Q" 11;
    Hashtbl.add card_value "J" 10;
    Hashtbl.add card_value "T" 9;
    Hashtbl.add card_value "9" 8;
    Hashtbl.add card_value "8" 7;
    Hashtbl.add card_value "7" 6;
    Hashtbl.add card_value "6" 5;
    Hashtbl.add card_value "5" 4;
    Hashtbl.add card_value "4" 3;
    Hashtbl.add card_value "3" 2;
    Hashtbl.add card_value "2" 1;
    Hashtbl.add card_value "" (-1);
    

    let count_cards {cards = c; points=_} = 
      let table = Hashtbl.create 5 in 
      let rec loop = function
        | [] -> () 
        | h::t -> (if Hashtbl.mem table h 
                  then Hashtbl.replace table h (Hashtbl.find table h + 1)     
                  else Hashtbl.add table h 1; loop t)
      in loop c;       
      let seq = Hashtbl.to_seq table in
      let res = List.of_seq seq in
      List.sort (fun (card1, value1) (card2, value2) -> 
        if value1 > value2 then -1 else if value2 > value1 then 1 
        else if Hashtbl.find card_value card1 > Hashtbl.find card_value card2 then -1 
        else if Hashtbl.find card_value card2 > Hashtbl.find card_value card1 then 1
        else 0) res  
    in let l1 = count_cards h1 and l2 = count_cards h2 in
    let rec cmp_val = function | [], [] -> 0  
      | (_, num1)::t1, (_, num2)::t2 -> 
      if num1 > num2 then 1 else if num2 > num1 then -1 
      else cmp_val (t1, t2) 
      | _ -> failwith "compare"
    in
    let rec cmp_c = function | [], [] -> 0 
      | card1::t1, card2::t2 -> 
      if Hashtbl.find card_value card1 > Hashtbl.find card_value card2 then 1
      else if Hashtbl.find card_value card2 > Hashtbl.find card_value card1 then -1 
      else cmp_c (t1, t2)
      | _ -> failwith "compare"
    in
    let cmp = cmp_val (l1, l2) in if cmp = 0 then cmp_c (h1.cards, h2.cards) else cmp


  let cmp_cards2 (h1 : hand) (h2 : hand) : int = 
    let card_value = Hashtbl.create 10 in 
    Hashtbl.add card_value "A" 13;
    Hashtbl.add card_value "K" 12;
    Hashtbl.add card_value "Q" 11;
    Hashtbl.add card_value "J" 0;
    Hashtbl.add card_value "T" 9;
    Hashtbl.add card_value "9" 8;
    Hashtbl.add card_value "8" 7;
    Hashtbl.add card_value "7" 6;
    Hashtbl.add card_value "6" 5;
    Hashtbl.add card_value "5" 4;
    Hashtbl.add card_value "4" 3;
    Hashtbl.add card_value "3" 2;
    Hashtbl.add card_value "2" 1;
    Hashtbl.add card_value "" (-1);
    

    let count_cards {cards = c; points=_} = 
      let table = Hashtbl.create 5 in 
      let nbJ = ref 0 in
      let rec loop = function
        | [] -> () 
        | h::t -> (if h = "J" then nbJ := !nbJ + 1 else if Hashtbl.mem table h 
                  then Hashtbl.replace table h (Hashtbl.find table h + 1)     
                  else Hashtbl.add table h 1; loop t)
      in loop c;       
      let seq = Hashtbl.to_seq table in
      let tmp = List.of_seq seq in
      let sorted = List.sort (fun (card1, value1) (card2, value2) -> 
        if value1 > value2 then -1 else if value2 > value1 then 1 
        else if Hashtbl.find card_value card1 > Hashtbl.find card_value card2 then -1 
        else if Hashtbl.find card_value card2 > Hashtbl.find card_value card1 then 1
        else 0) tmp in match sorted with | (s, n)::t -> (s, n + !nbJ)::t | [] -> [("J", 5)]  
    in let l1 = count_cards h1 and l2 = count_cards h2 in
    let rec cmp_val = function | [], [] -> 0  
      | (_, num1)::t1, (_, num2)::t2 -> 
      if num1 > num2 then 1 else if num2 > num1 then -1 
      else cmp_val (t1, t2) 
      | _ -> failwith "compare"
    in
    let rec cmp_c = function | [], [] -> 0 
      | card1::t1, card2::t2 -> 
      if Hashtbl.find card_value card1 > Hashtbl.find card_value card2 then 1
      else if Hashtbl.find card_value card2 > Hashtbl.find card_value card1 then -1 
      else cmp_c (t1, t2)
      | _ -> failwith "compare"
    in
    let cmp = cmp_val (l1, l2) in if cmp = 0 then cmp_c (h1.cards, h2.cards) else cmp


  let part1 file_name = 
    let list_line = match Pars.read_lines file_name with 
      | Some l -> l 
      | None -> failwith "parsing" in
    let aux_fold s = match Pars.split s " " with 
      | [c; n] -> {cards = Pars.split c ""; points = int_of_string n} 
      | _ -> failwith "parsing" 
    in
    let hands = List.fold_left (fun acc s -> (aux_fold s)::acc) [] list_line in
    let sorted = ListLabels.sort ~cmp:cmp_cards hands in
    
    let rec fold ind tmp = function 
      | [] -> (tmp) 
      | h::t -> (fold (ind+1) (tmp+ind*h.points) t)  
    in fold 1 0 sorted



  let part2 file_name = 
    let list_line = match Pars.read_lines file_name with 
      | Some l -> l 
      | None -> failwith "parsing" in
    let aux_fold s = match Pars.split s " " with 
      | [c; n] -> {cards = Pars.split c ""; points = int_of_string n} 
      | _ -> failwith "parsing" 
    in
    let hands = List.fold_left (fun acc s -> (aux_fold s)::acc) [] list_line in
    let sorted = ListLabels.sort ~cmp:cmp_cards2 hands in
    
    let rec fold ind tmp = function 
      | [] -> (tmp) 
      | h::t -> (fold (ind+1) (tmp+ind*h.points) t)  
    in fold 1 0 sorted

end 


let () =
  if Array.length Sys.argv <> 2 then begin
    Printf.printf "Usage: %s <filename>\n" Sys.argv.(0);
    exit 1;
  end else
    let file_name = Sys.argv.(1) in
    let r1 = Day7.part1 file_name in 
    let r2 = Day7.part2 file_name in 
    (print_int r1; print_newline (); print_int r2; print_newline ())    

