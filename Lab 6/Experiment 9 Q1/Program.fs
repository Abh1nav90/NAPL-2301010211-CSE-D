open System

let rec readMarks subject =
    printf "Enter marks for %s (0-100): " subject

    match Double.TryParse(Console.ReadLine()) with
    | true, mark when mark >= 0.0 && mark <= 100.0 -> mark
    | _ ->
        printfn "Please enter a valid mark between 0 and 100."
        readMarks subject

let classify percentage =
    match percentage with
    | value when value >= 60.0 -> "First Division"
    | value when value >= 50.0 -> "Second Division"
    | value when value >= 40.0 -> "Third Division"
    | _ -> "Fail"

[<EntryPoint>]
let main _ =
    printfn "===== Marks Percentage and Division ====="

    let subjectNames = [ "Subject 1"; "Subject 2"; "Subject 3"; "Subject 4" ]
    let marks = subjectNames |> List.map readMarks

    let total = List.sum marks
    let percentage = total / 4.0

    printfn "\nTotal marks: %.2f / 400" total
    printfn "Percentage: %.2f%%" percentage
    printfn "Division: %s" (classify percentage)

    0