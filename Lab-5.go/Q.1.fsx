let calculatePercentage marks =
    let total = List.sum marks
    float total / 4.0

let getDivision percentage =
    match percentage with
    | p when p >= 60.0 -> "First Division"
    | p when p >= 50.0 -> "Second Division"
    | p when p >= 40.0 -> "Third Division"
    | _ -> "Fail"

printf "Enter marks for Subject 1: "
let m1 = int (System.Console.ReadLine())

printf "Enter marks for Subject 2: "
let m2 = int (System.Console.ReadLine())

printf "Enter marks for Subject 3: "
let m3 = int (System.Console.ReadLine())

printf "Enter marks for Subject 4: "
let m4 = int (System.Console.ReadLine())

let marks = [m1; m2; m3; m4]
let percentage = calculatePercentage marks
let division = getDivision percentage

printfn "Percentage: %.2f%%" percentage
printfn "Division: %s" division