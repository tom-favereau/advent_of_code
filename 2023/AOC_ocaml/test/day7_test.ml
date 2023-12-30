open Alcotest


let tests = 
  [
    "test_part1", `Quick, (fun () -> assert (4 = 4)) 
  ]

let () =
  run "Day7 test" ["tests", tests]