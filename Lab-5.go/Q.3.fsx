open System
open System.Text.RegularExpressions

let checkPalindrome (s: string) =
    let reversed = s |> Seq.rev |> Seq.toArray |> String
    printfn "Original: %s | Reversed: %s | Palindrome: %b" s reversed (s = reversed)

let transformString (s: string) =
    if s.Length > 0 && (s |> Seq.forall Char.IsLetter) then
        s.ToUpper()
    elif s.Length > 0 && (s |> Seq.forall Char.IsDigit) then
        s + s
    else
        s

let isValidEmail (email: string) =
    Regex.IsMatch(email, @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$")

printfn "---- Palindrome Check ----"
["madam"; "hello"; "level"] |> List.iter checkPalindrome

printfn "\n---- String Transformation ----"
["hello"; "123"; "abc123"]
|> List.iter (fun s -> printfn "%s -> %s" s (transformString s))

printfn "\n---- Email Validation ----"
["user@gmail.com"; "abc@com"; "student@example.org"]
|> List.iter (fun email ->
    printfn "%s -> %b" email (isValidEmail email))