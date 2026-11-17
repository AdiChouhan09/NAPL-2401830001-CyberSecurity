namespace LibraryManager

open System

type LoanStatus =
    | Available
    | Borrowed

type Book =
    {
        Id: string
        Title: string
        Author: string
        Status: LoanStatus
    }

module Domain =
    let validate title author =
        if String.IsNullOrWhiteSpace title || String.IsNullOrWhiteSpace author then
            Error "Title and author are required."
        else
            Ok(title.Trim(), author.Trim())

    let statusText =
        function
        | Available -> "Available"
        | Borrowed -> "Borrowed"

    let available books =
        books |> List.filter (fun book -> book.Status = Available)

    let search (term: string) books =
        books
        |> List.filter (fun book ->
            book.Title.Contains(term, StringComparison.OrdinalIgnoreCase)
            || book.Author.Contains(term, StringComparison.OrdinalIgnoreCase))

    let format books =
        books
        |> List.map (fun book -> sprintf "%s | %s | %s | %s" book.Id book.Title book.Author (statusText book.Status))
