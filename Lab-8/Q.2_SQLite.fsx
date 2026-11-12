#r "nuget: Microsoft.Data.Sqlite"

open Microsoft.Data.Sqlite

// Connect to SQLite database
let connection = new SqliteConnection("Data Source=Students.db")

connection.Open()

// Create Students table
let createTable = connection.CreateCommand()

createTable.CommandText <-
    "CREATE TABLE IF NOT EXISTS Students
     (Id INTEGER PRIMARY KEY, Name TEXT, Age INTEGER)"

createTable.ExecuteNonQuery() |> ignore

printfn "Students table created successfully."

// Function to display student records
let displayStudents () =
    let command = connection.CreateCommand()
    command.CommandText <- "SELECT * FROM Students"

    use reader = command.ExecuteReader()

    if reader.HasRows then
        while reader.Read() do
            printfn "ID: %d, Name: %s, Age: %d" (reader.GetInt32(0)) (reader.GetString(1)) (reader.GetInt32(2))
    else
        printfn "No student records found."

// CREATE operation
let insert = connection.CreateCommand()

insert.CommandText <-
    "INSERT OR REPLACE INTO Students (Id, Name, Age)
     VALUES (1, 'Aditya', 19)"

insert.ExecuteNonQuery() |> ignore

printfn "\nCREATE: Student inserted successfully."

// READ operation
printfn "\nREAD: Student Records"

displayStudents ()

// UPDATE operation
let update = connection.CreateCommand()

update.CommandText <- "UPDATE Students SET Age = 20 WHERE Id = 1"

update.ExecuteNonQuery() |> ignore

printfn "\nUPDATE: Student age updated successfully."

displayStudents ()

// DELETE operation
let delete = connection.CreateCommand()

delete.CommandText <- "DELETE FROM Students WHERE Id = 1"

delete.ExecuteNonQuery() |> ignore

printfn "\nDELETE: Student deleted successfully."

displayStudents ()

// Close connection
connection.Close()

printfn "\nDatabase operations completed successfully."
