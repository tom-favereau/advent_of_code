module MyQueue : sig

  type 'a queue

  val empty : 'a queue
  
  val push : 'a -> 'a queue -> 'a queue

  val pop : 'a queue -> 'a option * 'a queue

  val length : 'a queue -> int


end = struct 
  type 'a queue = 'a list * 'a list
  
  let empty = ([], [])

  let push x (front, rear) = (front, x::rear)

  let pop = function 
    | [], [] -> (None, ([], []))
    | [], rear -> let tmp = List.rev rear in (Some (List.hd tmp), (List.tl tmp, []))
    | h::t, rear -> (Some h, (t, rear))  
  
  let length (front, rear) = 
    (List.length front) + (List.length rear) 
end 
  

module Pars : sig 


  val read_lines : string -> string list option 

  val read : string -> string option

  val split : string -> string -> string list

end = struct 



  let split s exp = Str.split (Str.regexp exp) s;;

  let read_lines file_name =
    try
      let channel = open_in file_name in
      let rec read_aux tmp =
        try
          let line = input_line channel in
          read_aux (line::tmp)
        with
        | End_of_file -> (close_in channel; tmp)
      in
      Some (read_aux [])
    with
    | Sys_error err ->( print_endline ("Error: " ^ err); None)


  let read file_name = 
    try 
      let ic = open_in file_name in
      let len = in_channel_length ic in
      let content = really_input_string ic len in 
      close_in ic;
      Some content
    with 
    | Sys_error err -> (print_endline ("Error: " ^ err); None)
end 

module Debug : sig 
  val show_int : int -> unit 

  val show_float : float -> unit 

  val show_list : 'a list -> ('a -> unit) -> unit

  val show_array : 'a array -> ('a -> unit) -> unit 
  
end = struct 

  let show_int x = Printf.printf "%d" x

  let show_float x = Printf.printf "%f" x

  let show_list l show_element = List.iter show_element l

  let show_array arr show_element = Array.iter show_element arr 

end
