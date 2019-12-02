let fuel1 acc value =
  acc + (value / 3) - 2

let rec fuel2 acc value =
  let value = (value / 3) - 2 in
  if value > 0
  then fuel2 (acc + value) value
  else acc

let rec main input acc1 acc2 =
  match input with
  | [] ->
    print_int acc1;
    print_string " ";
    print_int acc2;
    print_string "\n"
  | e::l ->
    let value = int_of_string e in
    main l (fuel1 acc1 value) (fuel2 acc2 value)

let read_lines name : string list =
  let ic = open_in name in
  let try_read () =
    try Some (input_line ic) with End_of_file -> None in
  let rec loop acc = match try_read () with
    | Some s -> loop (s :: acc)
    | None -> close_in ic; List.rev acc in
  loop []

let _ =
  let input = read_lines "input.txt" in
  main input 0 0
