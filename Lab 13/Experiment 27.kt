class Student(
    rollNo: Int,
    name: String,
    course: String,
    percentage: Double
) {
    // Public getters and private setters demonstrate controlled access.
    var rollNo: Int = rollNo
        private set
    var name: String = name
        private set
    var course: String = course
        private set
    var percentage: Double = percentage
        private set

    // Member function 1: accept and update student information from input.
    fun updateFromInput() {
        println("\nEnter updated details for student:")

        print("Roll Number: ")
        val newRollNo = readLine()?.trim()?.toIntOrNull()
        if (newRollNo != null) rollNo = newRollNo
        else println("Invalid roll number; previous value retained.")

        print("Name: ")
        val newName = readLine()?.trim()
        if (!newName.isNullOrBlank()) name = newName
        else println("Name cannot be blank; previous value retained.")

        print("Course: ")
        val newCourse = readLine()?.trim()
        if (!newCourse.isNullOrBlank()) course = newCourse
        else println("Course cannot be blank; previous value retained.")

        print("Percentage: ")
        val newPercentage = readLine()?.trim()?.toDoubleOrNull()
        if (newPercentage != null && newPercentage in 0.0..100.0) {
            percentage = newPercentage
        } else {
            println("Enter a percentage from 0 to 100; previous value retained.")
        }
    }

    // Member function 2: display all details in a formatted manner.
    fun displayDetails() {
        println("----------------------------------------")
        println("             STUDENT DETAILS")
        println("----------------------------------------")
        println("Roll Number : $rollNo")
        println("Name        : $name")
        println("Course      : $course")
        println("Percentage  : %.2f%%".format(percentage))
        println("----------------------------------------")
    }
}

fun main() {
    // Create at least two Student objects using the primary constructor.
    val student1 = Student(101, "Abhinav", "CSE", 85.5)
    val student2 = Student(102, "Rahul", "CSE", 88.0)

    println("Initial Student Records")
    student1.displayDetails()
    student2.displayDetails()

    // Call both member functions on each object.
    println("\nUpdate Student 1")
    student1.updateFromInput()
    student1.displayDetails()

    println("\nUpdate Student 2")
    student2.updateFromInput()
    student2.displayDetails()
}