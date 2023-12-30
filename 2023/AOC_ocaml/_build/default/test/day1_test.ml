(* Include your source file *)
(*include Day1*)

(* Define your test cases *)
let%test_exempleP1 _ =
  (* Your test logic goes here *)
  let r, _ = pars "inputs/input_test01.txt" in
  assert (r  = 0)
