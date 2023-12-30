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
  
