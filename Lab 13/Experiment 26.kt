// 1. Function with two default parameters.
fun displayEmployee(name: String = "Unknown", department: String = "General") {
    println("Employee Name: $name")
    println("Department: $department")
}

// 3. Function using vararg to accept a variable number of Int arguments.
fun printNumbers(vararg numbers: Int) {
    println("Numbers received:")
    for (number in numbers) {
        println(number)
    }
}

// 4. Higher-order function accepting a lambda (Int) -> Int.
fun applyOperation(numbers: List<Int>, operation: (Int) -> Int): List<Int> {
    return numbers.map(operation)
}

fun main() {
    println("--- Default Arguments ---")
    displayEmployee() // Call without supplying either default argument.
    println()
    displayEmployee("Abhinav", "CSE") // Call with both arguments.
    println()

    println("--- Named Arguments ---")
    // Arguments are passed in an order different from the declaration.
    displayEmployee(department = "Information Technology", name = "Rahul")
    println()

    println("--- Vararg Function ---")
    printNumbers(10, 20, 30, 40, 50)
    println()

    println("--- Higher-Order Function ---")
    val values = listOf(1, 2, 3, 4, 5)
    println("Original List: $values")
    val squares = applyOperation(values) { number -> number * number }
    println("Squares: $squares")
    val doubled = applyOperation(values) { number -> number * 2 }
    println("Doubled Values: $doubled")
}