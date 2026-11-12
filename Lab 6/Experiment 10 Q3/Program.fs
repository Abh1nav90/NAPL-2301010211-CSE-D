open System
open System.Numerics
open System.Text.RegularExpressions

let reverseText (value: string) =
    value |> Seq.rev |> Array.ofSeq |> fun chars -> String(chars)

let checkPalindrome (value: string) =
    String.Equals(value, reverseText value, StringComparison.OrdinalIgnoreCase)

let changeText (value: string) =
    match value with
    | text when Regex.IsMatch(text, @"^[A-Za-z]+$") ->
        text.ToUpperInvariant()
    | text when Regex.IsMatch(text, @"^[0-9]+$") ->
        (BigInteger.Parse(text) * 2I).ToString()
    | _ -> value

let checkEmail (address: string) =
    let emailPattern =
        @"^[A-Za-z0-9.!#$%&'*+/=?^_`{|}~-]+@[A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?(?:\.[A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?)+$"

    Regex.IsMatch(address, emailPattern)

[<EntryPoint>]
let main _ =
    printfn "===== 1. Reverse a string and check palindrome ====="

    [ "radar"; "world"; "Level" ]
    |> List.iter (fun item ->
        printfn "Input: %s | Reversed: %s | Palindrome: %b"
            item (reverseText item) (checkPalindrome item))

    printfn "\n===== 2. Transform a string ====="

    [ "functional"; "125"; "abc123" ]
    |> List.iter (fun item ->
        printfn "Input: %s | Result: %s" item (changeText item))

    printfn "\n===== 3. Validate email addresses ====="

    [ "2301010211@krmu.edu.in"; "abhinav.jha@gmail.com"; "wrong-address" ]
    |> List.iter (fun address ->
        printfn "%s -> %s"
            address
            (if checkEmail address then "Valid" else "Invalid"))

    0
