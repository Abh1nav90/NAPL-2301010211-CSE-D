open System
open Microsoft.Data.Sqlite

// Student domain model
type Student = {
    Id: int
    Name: string
    Department: string
    Marks: float
}

// SQLite database configuration
let connectionString = "Data Source=students.db"

// Open and use a database connection
let withConnection (work: SqliteConnection -> Async<'T>) =
    async {
        use connection = new SqliteConnection(connectionString)
        do! connection.OpenAsync() |> Async.AwaitTask
        return! work connection
    }

// Create the student table
let initializeDatabase () =
    withConnection (fun connection ->
        async {
            use command = connection.CreateCommand()

            command.CommandText <-
                """
                CREATE TABLE IF NOT EXISTS Students (
                    Id INTEGER PRIMARY KEY,
                    Name TEXT NOT NULL,
                    Department TEXT NOT NULL,
                    Marks REAL NOT NULL
                )
                """

            let! _ =
                command.ExecuteNonQueryAsync()
                |> Async.AwaitTask

            printfn "Database initialized successfully."
        }
    )

let addStudent student =
    withConnection (fun connection ->
        async {
            use command = connection.CreateCommand()

            command.CommandText <-
                """
                INSERT INTO Students (Id, Name, Department, Marks)
                VALUES (@id, @name, @department, @marks)
                """

            command.Parameters.AddWithValue("@id", student.Id) |> ignore
            command.Parameters.AddWithValue("@name", student.Name) |> ignore
            command.Parameters.AddWithValue("@department", student.Department) |> ignore
            command.Parameters.AddWithValue("@marks", student.Marks) |> ignore

            let! result =
                command.ExecuteNonQueryAsync()
                |> Async.AwaitTask

            return result > 0
        }
    )

let getStudents () =
    withConnection (fun connection ->
        async {
            use command = connection.CreateCommand()

            command.CommandText <-
                "SELECT Id, Name, Department, Marks FROM Students ORDER BY Id"

            use! reader =
                command.ExecuteReaderAsync()
                |> Async.AwaitTask

            let students = ResizeArray<Student>()

            while reader.Read() do
                students.Add({
                    Id = reader.GetInt32(0)
                    Name = reader.GetString(1)
                    Department = reader.GetString(2)
                    Marks = reader.GetDouble(3)
                })

            return List.ofSeq students
        }
    )

let updateStudent id name department marks =
    withConnection (fun connection ->
        async {
            use command = connection.CreateCommand()

            command.CommandText <-
                """
                UPDATE Students
                SET Name = @name,
                    Department = @department,
                    Marks = @marks
                WHERE Id = @id
                """

            command.Parameters.AddWithValue("@id", id) |> ignore
            command.Parameters.AddWithValue("@name", name) |> ignore
            command.Parameters.AddWithValue("@department", department) |> ignore
            command.Parameters.AddWithValue("@marks", marks) |> ignore

            let! result =
                command.ExecuteNonQueryAsync()
                |> Async.AwaitTask

            return result > 0
        }
    )

let deleteStudent id =
    withConnection (fun connection ->
        async {
            use command = connection.CreateCommand()

            command.CommandText <-
                "DELETE FROM Students WHERE Id = @id"

            command.Parameters.AddWithValue("@id", id) |> ignore

            let! result =
                command.ExecuteNonQueryAsync()
                |> Async.AwaitTask

            return result > 0
        }
    )

// Pattern matching for result classification
let classifyMarks marks =
    match marks with
    | value when value >= 60.0 -> "First Division"
    | value when value >= 50.0 -> "Second Division"
    | value when value >= 40.0 -> "Third Division"
    | _ -> "Fail"

// Read validated integer input
let readInteger message =
    let rec read () =
        printf "%s" message

        match Int32.TryParse(Console.ReadLine()) with
        | true, value when value > 0 -> value
        | _ ->
            printfn "Please enter a valid positive integer."
            read ()

    read ()

// Read validated marks
let readMarks () =
    let rec read () =
        printf "Enter marks (0-100): "

        match Double.TryParse(Console.ReadLine()) with
        | true, value when
            not (Double.IsNaN(value))
            && not (Double.IsInfinity(value))
            && value >= 0.0
            && value <= 100.0 -> value
        | _ ->
            printfn "Enter valid marks between 0 and 100."
            read ()

    read ()

// Display student information
let displayStudents students =
    if List.isEmpty students then
        printfn "No student records found."
    else
        printfn "\n========== STUDENT RECORDS =========="

        students
        |> List.iter (fun student ->
            printfn "ID: %d | Name: %s | Department: %s | Marks: %.2f | Result: %s"
                student.Id
                student.Name
                student.Department
                student.Marks
                (classifyMarks student.Marks)
        )

// Functional processing using filter and map
let showPassedStudents students =
    let results =
        students
        |> List.filter (fun student -> student.Marks >= 40.0)
        |> List.map (fun student ->
            student.Name, student.Marks, classifyMarks student.Marks
        )

    printfn "\n========== PASSED STUDENTS =========="

    if List.isEmpty results then
        printfn "No passed students found."
    else
        results
        |> List.iter (fun (name, marks, division) ->
            printfn "%s - %.2f - %s" name marks division
        )

// Menu-driven application
let rec menuLoop () =
    async {
        printfn "\n================================"
        printfn "     STUDENT MANAGEMENT SYSTEM"
        printfn "================================"
        printfn "1. Add Student"
        printfn "2. Display All Students"
        printfn "3. Update Student"
        printfn "4. Delete Student"
        printfn "5. Show Passed Students"
        printfn "0. Exit"
        printfn "================================"
        printf "Choose an option: "

        let choice = Console.ReadLine()

        match choice with
        | "1" ->
            let id = readInteger "Enter student ID: "

            printf "Enter student name: "
            let name = Console.ReadLine()

            printf "Enter department: "
            let department = Console.ReadLine()

            let marks = readMarks ()

            if String.IsNullOrWhiteSpace(name)
               || String.IsNullOrWhiteSpace(department) then
                printfn "Name and department cannot be empty."
            else
                let student = {
                    Id = id
                    Name = name.Trim()
                    Department = department.Trim()
                    Marks = marks
                }

                try
                    let! added = addStudent student

                    if added then
                        printfn "Student added successfully."
                    else
                        printfn "Student could not be added."
                with ex ->
                    printfn "Database error: %s" ex.Message

        | "2" ->
            try
                let! students = getStudents ()
                displayStudents students
            with ex ->
                printfn "Database error: %s" ex.Message

        | "3" ->
            let id = readInteger "Enter student ID to update: "

            printf "Enter new name: "
            let name = Console.ReadLine()

            printf "Enter new department: "
            let department = Console.ReadLine()

            let marks = readMarks ()

            if String.IsNullOrWhiteSpace(name)
               || String.IsNullOrWhiteSpace(department) then
                printfn "Name and department cannot be empty."
            else
                try
                    let! updated =
                        updateStudent id name department marks

                    if updated then
                        printfn "Student updated successfully."
                    else
                        printfn "Student ID not found."
                with ex ->
                    printfn "Database error: %s" ex.Message

        | "4" ->
            let id = readInteger "Enter student ID to delete: "

            try
                let! deleted = deleteStudent id

                if deleted then
                    printfn "Student deleted successfully."
                else
                    printfn "Student ID not found."
            with ex ->
                printfn "Database error: %s" ex.Message

        | "5" ->
            try
                let! students = getStudents ()
                showPassedStudents students
            with ex ->
                printfn "Database error: %s" ex.Message

        | "0" ->
            printfn "Student Management System closed."

        | _ ->
            printfn "Invalid choice. Please try again."

        if choice <> "0" then
            do! menuLoop ()
    }

[<EntryPoint>]
let main _ =
    try
        async {
            do! initializeDatabase ()
            do! menuLoop ()
        }
        |> Async.RunSynchronously

        0
    with ex ->
        printfn "Application error: %s" ex.Message
        1