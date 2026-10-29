let isPrime n =
    if n < 2 then false
    else
        let rec check divisor =
            if divisor * divisor > n then true
            elif n % divisor = 0 then false
            else check (divisor + 1)
        check 2

let numbers = [7; 10; 13]

numbers
|> List.iter (fun n ->
    if isPrime n then
        printfn "%d is a prime number" n
    else
        printfn "%d is not a prime number" n)