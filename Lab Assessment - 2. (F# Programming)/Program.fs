namespace LibraryManager

open System
open MongoDB.Bson

module Program =
    let prompt label =
        printf "%s: " label
        match Console.ReadLine() with null -> "" | value -> value

    let show books =
        match Domain.format books with
        | [] -> printfn "No books found."
        | rows -> rows |> List.iter (printfn "%s")

    let report = function Ok () -> printfn "Saved." | Error message -> printfn "%s" message

    let readId () =
        match ObjectId.TryParse(prompt "Book ID") with
        | true, id -> Some (id.ToString())
        | _ -> printfn "Enter the 24-character MongoDB book ID."; None

    let runMenu (repository: Repository) = async {
        let mutable running = true
        while running do
            printfn "\n1 Add  2 List  3 Update  4 Delete  5 Available  6 Search  0 Exit"
            let choice = Console.ReadLine()
            try
                match choice with
                | null | "0" -> running <- false
                | "1" ->
                    let title = prompt "Title"
                    let author = prompt "Author"
                    let! result = repository.Create(title, author)
                    report result
                | "2" | "5" | "6" ->
                    let! books = repository.ReadAll()
                    match choice with
                    | "5" -> books |> Domain.available |> show
                    | "6" -> books |> Domain.search (prompt "Search title or author") |> show
                    | _ -> show books
                | "3" ->
                    match readId () with
                    | None -> ()
                    | Some id ->
                        let title = prompt "New title"
                        let author = prompt "New author"
                        match (prompt "Status (a = available, b = borrowed)").ToLowerInvariant() with
                        | "a" | "b" as code ->
                            let status = if code = "a" then Available else Borrowed
                            let! result = repository.Update(id, title, author, status)
                            report result
                        | _ -> printfn "Invalid status. No change saved."
                | "4" ->
                    match readId () with
                    | None -> ()
                    | Some id ->
                        let! result = repository.Delete id
                        report result
                | _ -> printfn "Choose a listed option."
            with ex -> printfn "Operation failed: %s" ex.Message
    }

    let setting name fallback =
        match Environment.GetEnvironmentVariable(name) with
        | null | "" -> fallback
        | value -> value

    let selfTest () = async {
        let uri = setting "MONGODB_URI" "mongodb://localhost:27017"
        let database = "library_test_" + Guid.NewGuid().ToString("N")
        let check condition message = if not condition then failwith message
        try
            use repo = new Repository(uri, database)
            do! repo.Initialize()
            let! invalid = repo.Create(" ", "Author")
            check (Result.isError invalid) "Validation failed"
            let! created = repo.Create("F# in Action", "O'Brien")
            check (created = Ok ()) "Create failed"
            use fresh = new Repository(uri, database)
            let! books = fresh.ReadAll()
            check (books.Length = 1 && books.Head.Author = "O'Brien") "Persistence failed"
            check ((Domain.search "f#" books).Length = 1) "Search failed"
            check ((Domain.available books).Length = 1) "Filter failed"
            check ((Domain.format books).Length = 1) "Map failed"
            let id = books.Head.Id
            let! updated = repo.Update(id, "Updated", "Author", Borrowed)
            check (updated = Ok ()) "Update failed"
            let! changed = repo.ReadAll()
            check (changed.Head.Title = "Updated" && changed.Head.Status = Borrowed) "Update persistence failed"
            check (Domain.available changed |> List.isEmpty) "Borrowed filter failed"
            let! deleted = repo.Delete id
            check (deleted = Ok ()) "Delete failed"
            let! missing = repo.Delete id
            check (Result.isError missing) "Missing ID failed"
            let! empty = repo.ReadAll()
            check empty.IsEmpty "Delete persistence failed"
            printfn "PASS: validation, CRUD, persistence, map, filter, search, missing ID."
        finally
            use cleanupClient = new MongoDB.Driver.MongoClient(uri)
            cleanupClient.DropDatabase(database)
    }

    [<EntryPoint>]
    let main args =
        try
            if args = [| "--self-test" |] then selfTest () |> Async.RunSynchronously
            else
                let uri = setting "MONGODB_URI" "mongodb://localhost:27017"
                let database = setting "MONGODB_DATABASE" "fsharp_library"
                printfn "Library Manager - MongoDB database: %s" database
                use repository = new Repository(uri, database)
                async {
                    do! repository.Initialize()
                    do! runMenu repository
                } |> Async.RunSynchronously
            0
        with ex ->
            eprintfn "Error: %s" ex.Message
            1
