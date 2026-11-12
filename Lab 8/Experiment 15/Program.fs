open System
open Microsoft.Data.Sqlite
open MongoDB.Bson
open MongoDB.Driver

// -------------------- SQLite helpers --------------------

let showSqliteStudents (connection: SqliteConnection) =
    use command = connection.CreateCommand()
    command.CommandText <- "SELECT Id, Name, Age FROM Students ORDER BY Id"

    use reader = command.ExecuteReader()
    let mutable foundRecord = false

    printfn "SQLite records:"
    while reader.Read() do
        foundRecord <- true
        printfn "  Id: %d | Name: %s | Age: %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

    if not foundRecord then
        printfn "  (no records)"

let runSqliteDemo () =
    printfn "\n========== RELATIONAL DATABASE: SQLITE =========="

    use connection = new SqliteConnection("Data Source=students.db")
    connection.Open()

    use setup = connection.CreateCommand()
    setup.CommandText <-
        "CREATE TABLE IF NOT EXISTS Students (Id INTEGER PRIMARY KEY, Name TEXT NOT NULL, Age INTEGER NOT NULL)"
    setup.ExecuteNonQuery() |> ignore

    // Use a known ID so the demonstration can be run more than once.
    let studentId = 204

    use removePrevious = connection.CreateCommand()
    removePrevious.CommandText <- "DELETE FROM Students WHERE Id = $id"
    removePrevious.Parameters.AddWithValue("$id", studentId) |> ignore
    removePrevious.ExecuteNonQuery() |> ignore

    // CREATE
    use addStudent = connection.CreateCommand()
    addStudent.CommandText <-
        "INSERT INTO Students (Id, Name, Age) VALUES ($id, $name, $age)"
    addStudent.Parameters.AddWithValue("$id", studentId) |> ignore
    addStudent.Parameters.AddWithValue("$name", "Mira") |> ignore
    addStudent.Parameters.AddWithValue("$age", 20) |> ignore
    addStudent.ExecuteNonQuery() |> ignore
    printfn "\n1. CREATE - inserted a student"
    showSqliteStudents connection

    // READ
    use readStudent = connection.CreateCommand()
    readStudent.CommandText <- "SELECT Id, Name, Age FROM Students WHERE Id = $id"
    readStudent.Parameters.AddWithValue("$id", studentId) |> ignore
    use reader = readStudent.ExecuteReader()
    printfn "\n2. READ - selected the inserted student"
    if reader.Read() then
        printfn "  Id: %d | Name: %s | Age: %d"
            (reader.GetInt32(0)) (reader.GetString(1)) (reader.GetInt32(2))
    else
        printfn "  Student was not found."
    reader.Close()
    showSqliteStudents connection

    // UPDATE
    use changeStudent = connection.CreateCommand()
    changeStudent.CommandText <-
        "UPDATE Students SET Name = $name, Age = $age WHERE Id = $id"
    changeStudent.Parameters.AddWithValue("$name", "Mira Sharma") |> ignore
    changeStudent.Parameters.AddWithValue("$age", 21) |> ignore
    changeStudent.Parameters.AddWithValue("$id", studentId) |> ignore
    changeStudent.ExecuteNonQuery() |> ignore
    printfn "\n3. UPDATE - changed the student's details"
    showSqliteStudents connection

    // DELETE
    use deleteStudent = connection.CreateCommand()
    deleteStudent.CommandText <- "DELETE FROM Students WHERE Id = $id"
    deleteStudent.Parameters.AddWithValue("$id", studentId) |> ignore
    deleteStudent.ExecuteNonQuery() |> ignore
    printfn "\n4. DELETE - removed the student"
    showSqliteStudents connection

// -------------------- MongoDB helpers --------------------

let showMongoStudents (students: IMongoCollection<BsonDocument>) =
    let records = students.Find(FilterDefinition<BsonDocument>.Empty).ToList()

    printfn "MongoDB records:"
    if records.Count = 0 then
        printfn "  (no records)"
    else
        records |> Seq.iter (fun record -> printfn "  %s" (record.ToJson()))

let runMongoDemo () =
    printfn "\n========== DOCUMENT DATABASE: MONGODB =========="

    // MongoDB must be running locally for this connection to succeed.
    let client = MongoClient("mongodb://localhost:27017")
    let database = client.GetDatabase("FSharpLabDatabase")
    let students = database.GetCollection<BsonDocument>("StudentRecords")

    let studentId = 204
    let byId = Builders<BsonDocument>.Filter.Eq("_id", BsonInt32(studentId))

    // Clear this sample ID first, allowing repeatable runs.
    students.DeleteOne(byId) |> ignore

    // CREATE
    let newStudent = BsonDocument()
    newStudent.Add("_id", BsonInt32(studentId))
    newStudent.Add("name", BsonString("Mira"))
    newStudent.Add("age", BsonInt32(20))
    students.InsertOne(newStudent)
    printfn "\n1. CREATE - inserted a student document"
    showMongoStudents students

    // READ
    printfn "\n2. READ - searching by student ID"
    let selected = students.Find(byId).FirstOrDefault()
    if isNull selected then
        printfn "  Student was not found."
    else
        printfn "  %s" (selected.ToJson())
    showMongoStudents students

    // UPDATE
    let changes =
        Builders<BsonDocument>.Update
            .Set("name", BsonString("Mira Sharma"))
            .Set("age", BsonInt32(21))

    students.UpdateOne(byId, changes) |> ignore
    printfn "\n3. UPDATE - changed the student's document"
    showMongoStudents students

    // DELETE
    students.DeleteOne(byId) |> ignore
    printfn "\n4. DELETE - removed the student document"
    showMongoStudents students

[<EntryPoint>]
let main _ =
    try
        runSqliteDemo ()
        runMongoDemo ()
        printfn "\nThe SQLite and MongoDB CRUD demonstrations have finished."
        0
    with error ->
        eprintfn "The database demonstration could not finish: %s" error.Message
        eprintfn "Check that the required packages are restored and MongoDB is running."
        1