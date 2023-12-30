open AOC_ocaml_lib.Utils

let split s exp = Str.split (Str.regexp exp) s;;



let pars file_name =
  try
    let channel = open_in file_name in

    let rec read_lines graphe =
      try
        let line = input_line channel in
        let liste = (split line ": ") in

        let sommet = List.hd liste in
        let voisins = split (List.hd (List.tl liste)) " " in

        let graphe =
          List.fold_left
            (fun graphe voisin ->
              let voisin = String.trim voisin in
              if Hashtbl.mem graphe sommet then
                Hashtbl.replace graphe sommet (voisin :: Hashtbl.find graphe sommet)
              else
                Hashtbl.add graphe sommet [voisin];
              if Hashtbl.mem graphe voisin then
                Hashtbl.replace graphe voisin (sommet :: Hashtbl.find graphe voisin)
              else
                Hashtbl.add graphe voisin [sommet];
              graphe)
            graphe voisins
        in
        read_lines graphe
      with
      | End_of_file -> (close_in channel; graphe)
    in

    let graphe = read_lines (Hashtbl.create 10) in
    graphe

  with
  | Sys_error err ->
    print_endline ("Error: " ^ err);
    Hashtbl.create 0
;;

let build_matrix graphe =
  let sommets = Hashtbl.length graphe in
  let matrice = Array.make_matrix sommets sommets 0 in
  let indexes = Hashtbl.create sommets in
  let ind = ref 0 in

  Hashtbl.iter
    (fun sommet voisins -> 
      let sommet_index = if Hashtbl.mem indexes sommet then Hashtbl.find indexes sommet 
                          else (Hashtbl.add indexes sommet !ind; ind := !ind+1; !ind-1) in 
      List.iter
        (fun voisin ->
          let voisin_index = if Hashtbl.mem indexes voisin then Hashtbl.find indexes voisin 
                              else (Hashtbl.add indexes voisin !ind; ind := !ind+1; !ind-1) in
          matrice.(sommet_index).(voisin_index) <- 1;
          matrice.(voisin_index).(sommet_index) <- 1; 
         )
        voisins
        )
    graphe;

  matrice
;;

let solve graphe = 
  let n = Array.length graphe in 
  let freq = Array.make_matrix n n 0 in
  let find_freq start =
    let visited = Array.make n false in
    let prev = Array.make n 0 in
    let queue = ref MyQueue.empty in
    queue := MyQueue.push start !queue;
    while MyQueue.length !queue > 0 do
      let node, n_queue = match MyQueue.pop !queue with | Some n, nq -> (n, nq) | _, nq -> 0, nq  in
      queue := n_queue;
      for i = 0 to n-1 do if graphe.(node).(i) = 1 && not visited.(i) then
        (visited.(i) <- true;
        queue := MyQueue.push i !queue;
        prev.(i) <- node)
      done;
    done;

    for i = 0 to n-1 do let node = ref i in 
      while !node <> start do let tmp = prev.(!node) in
        (freq.(min tmp !node).(max tmp !node) <- freq.(min tmp !node).(max tmp !node) + 1;
        node := tmp)
      done;
    done;
  in 
  for i = 5 to n-1 do find_freq i done;

  let m1 = ref 0 and e1 = ref (0, 0) in 
  let m2 = ref 0 and e2 = ref (0, 0) in 
  let m3 = ref 0 and e3 = ref (0, 0) in
  for i = 0 to n-1 do for j = 0 to n-1 do
    let v = freq.(i).(j) in
    if v > !m1 then (m3 := !m2; e3 := !e2; m2 := !m1; e2 := !e1; m1 := v; e1 := (i, j))
    else if v > !m2 then (m3 := !m2; e3 := !e2; m2 := v; e2 := (i, j))
    else if v > !m3 then (m3 := v; e3 := (i, j))
  done; done;

  let new_graphe = Array.make_matrix n n 0 in
  for i = 0 to n-1 do for j = 0 to n-1 do 
    new_graphe.(i).(j) <- graphe.(i).(j) 
  done; done; 
  let a, b = !e1 and c, d = !e2 and e, f = !e3 in
  let bfs = 
    let visited = Array.make n false in
    let queue = ref MyQueue.empty in 
    let number = ref 0 in
    queue := MyQueue.push 0 !queue;
    while MyQueue.length !queue > 0 do 
      let node, new_queue = match MyQueue.pop !queue with | Some n, nq -> n, nq | _, nq -> 0, nq in
      queue := new_queue;
      for i = 0 to n-1 do 
        if not visited.(i) && graphe.(node).(i) = 1 then 
          (visited.(i) <- true; queue := MyQueue.push i !queue; number := !number+1)
      done;
    done;
    (!number, n - !number)
  in 
  ( 
  new_graphe.(a).(b) <- 0; new_graphe.(c).(d) <- 0; new_graphe.(e).(f) <- 0;
  print_int a; print_string " "; print_int b; print_newline (); 
  print_int c; print_string " "; print_int d; print_newline ();
  print_int e; print_string " "; print_int f; print_newline ();
  for i = 0 to n-1 do for j = 0 to n-1 do (print_int freq.(i).(j); print_string " ") done; print_newline (); done; 
  bfs 
  )   
;;

(*

let afficher_graphe graphe =
  Hashtbl.iter
    (fun sommet voisins ->
      let voisins_str = String.concat " " voisins in
      Printf.printf "%s: %s\n" sommet voisins_str)
    graphe
;;



let print_matrix m = 
  for i = 0 to Array.length m -1 do for j = 0 to Array.length m.(0) -1 do 
    print_int m.(i).(j) done; print_endline ""; done; print_newline ();
;;
*)

let () =
  let fichier = "inputs/input_test25.txt" in
  let graphe = build_matrix (pars fichier) in
  let c1, c2 = solve graphe in   
  (print_int c1; print_string " "; print_int c2) 
;;



