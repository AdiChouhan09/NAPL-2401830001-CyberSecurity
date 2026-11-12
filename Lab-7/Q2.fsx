let list1 = [ 1; 2; 3; 4; 5 ]
let list2 = [ 6; 7; 8; 9; 10 ]

printfn "Original Lists:"
printfn "List 1: %A" list1
printfn "List 2: %A" list2

let mapped1 = List.map (fun x -> x * x) list1
let mapped2 = List.map (fun x -> x * x) list2

printfn "\nList.map - Square of elements:"
printfn "List 1: %A" mapped1
printfn "List 2: %A" mapped2

let filtered1 = List.filter (fun x -> x % 2 = 0) list1
let filtered2 = List.filter (fun x -> x % 2 = 0) list2

printfn "\nList.filter - Even numbers:"
printfn "List 1: %A" filtered1
printfn "List 2: %A" filtered2

let sum1 = List.fold (fun acc x -> acc + x) 0 list1
let sum2 = List.fold (fun acc x -> acc + x) 0 list2

printfn "\nList.fold - Sum of elements:"
printfn "List 1 Sum: %d" sum1
printfn "List 2 Sum: %d" sum2
