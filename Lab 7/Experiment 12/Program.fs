open System

[<EntryPoint>]
let main argv =

    // Two different input lists
    let list1 = [1; 2; 3; 4; 5]
    let list2 = [30; 35; 40; 45; 50]

    // 1. List.map - Square each element
    printfn "===== LIST.MAP ====="

    let square1 = List.map (fun x -> x * x) list1
    let square2 = List.map (fun x -> x * x) list2

    printfn "Original List 1: %A" list1
    printfn "Squared List 1: %A" square1

    printfn "Original List 2: %A" list2
    printfn "Squared List 2: %A" square2

    // 2. List.filter - Select even numbers
    printfn "\n===== LIST.FILTER ====="

    let even1 = List.filter (fun x -> x % 2 = 0) list1
    let even2 = List.filter (fun x -> x % 2 = 0) list2

    printfn "Original List 1: %A" list1
    printfn "Even Numbers: %A" even1

    printfn "Original List 2: %A" list2
    printfn "Even Numbers: %A" even2

    // 3. List.fold - Calculate sum
    printfn "\n===== LIST.FOLD ====="

    let sum1 = List.fold (fun acc x -> acc + x) 0 list1
    let sum2 = List.fold (fun acc x -> acc + x) 0 list2

    printfn "Sum of List 1: %d" sum1
    printfn "Sum of List 2: %d" sum2

    // Calculate product using List.fold
    let product1 = List.fold (fun acc x -> acc * x) 1 list1
    let product2 = List.fold (fun acc x -> acc * x) 1 list2

    printfn "Product of List 1: %d" product1
    printfn "Product of List 2: %d" product2

    0