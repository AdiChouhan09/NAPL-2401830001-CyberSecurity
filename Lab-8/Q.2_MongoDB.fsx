#r "nuget: MongoDB.Driver"

open MongoDB.Driver

// Student class
[<AllowNullLiteral>]
type Student() =
    member val Id: int = 0 with get, set
    member val Name: string = "" with get, set
    member val Age: int = 0 with get, set

// Connect to MongoDB
let client = new MongoClient("mongodb://localhost:27017")

let database = client.GetDatabase("Lab8")

let collection = database.GetCollection<Student>("Students")

// Filter for student record
let filter = Builders<Student>.Filter.Eq("Id", 1001)

// Remove previous test record if it exists
collection.DeleteOne(filter) |> ignore

// CREATE operation
let student = Student()
student.Id <- 1001
student.Name <- "Aditya"
student.Age <- 19

collection.InsertOne(student)

printfn "CREATE: Student inserted successfully."

// READ operation
let result = collection.Find(filter).FirstOrDefault()

printfn "\nREAD:"

if not (isNull result) then
    printfn "ID: %d, Name: %s, Age: %d" result.Id result.Name result.Age

// UPDATE operation
let update = Builders<Student>.Update.Set("Age", 20)

collection.UpdateOne(filter, update) |> ignore

printfn "\nUPDATE: Student age updated successfully."

let updatedResult = collection.Find(filter).FirstOrDefault()

if not (isNull updatedResult) then
    printfn "ID: %d, Name: %s, Age: %d" updatedResult.Id updatedResult.Name updatedResult.Age

// DELETE operation
collection.DeleteOne(filter) |> ignore

printfn "\nDELETE: Student deleted successfully."

// Verify deletion
let finalResult = collection.Find(filter).FirstOrDefault()

if isNull finalResult then
    printfn "No student records found."
else
    printfn "Student record still exists."

printfn "\nMongoDB operations completed successfully."
