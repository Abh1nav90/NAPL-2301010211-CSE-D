open System

// Checks possible divisors recursively until a factor is found
// or the square root of the number has been passed.
let isPrime number =
    let rec hasNoDivisor candidate =
        if candidate > number / candidate then
            true
        elif number % candidate = 0 then
            false
        else
            hasNoDivisor (candidate + 1)

    if number < 2 then false
    else hasNoDivisor 2

[<EntryPoint>]
let main _ =
    printf "Enter an integer: "

    match Int32.TryParse(Console.ReadLine()) with
    | true, value ->
        let result =
            if isPrime value then "is a prime number"
            else "is not a prime number"

        printfn "%d %s." value result
    | _ ->
        printfn "Invalid input. Please enter a whole number."

    0