
open System

// Task 1
let task1 =
    async {
        printfn "Task 1 started"

        do! Async.Sleep(3000)

        printfn "Task 1 completed"

        return "Result from Task 1"
    }

// Task 2
let task2 =
    async {
        printfn "Task 2 started"

        do! Async.Sleep(1000)

        printfn "Task 2 completed"

        return "Result from Task 2"
    }

// Task 3
let task3 =
    async {
        printfn "Task 3 started"

        do! Async.Sleep(2000)

        printfn "Task 3 completed"

        return "Result from Task 3"
    }

[<EntryPoint>]
let main argv =

    printfn "Starting asynchronous workflows...\n"

    // Execute all three tasks concurrently
    let tasks = [ task1; task2; task3 ]

    let results =
        tasks
        |> Async.Parallel
        |> Async.RunSynchronously

    printfn "\nAll tasks completed!"
    printfn "\n===== FINAL RESULTS ====="

    results
    |> Array.iter (fun result ->
        printfn "%s" result
    )

    0
