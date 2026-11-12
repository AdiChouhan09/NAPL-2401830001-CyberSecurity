let task (name: string) (delay: int) =
    async {
        printfn "%s started" name
        do! Async.Sleep(delay)
        printfn "%s completed" name
        return name + " Result"
    }

let tasks = [| task "Task 1" 3000; task "Task 2" 2000; task "Task 3" 1000 |]

printfn "Starting parallel execution...\n"

let results = Async.Parallel tasks |> Async.RunSynchronously

printfn "\nAll tasks completed!"
printfn "Collected Results:"

results |> Array.iter (printfn "%s")
