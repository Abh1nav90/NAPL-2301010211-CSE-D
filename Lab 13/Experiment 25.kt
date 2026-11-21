fun main() {

    // 1. Hello Kotlin
    println("Hello Kotlin!!!")

    // 2. Employee Bonus Calculation
    println("\n--- Employee Bonus Calculation ---")

    print("Enter department code (IT / HR / FINANCE): ")
    val deptCode = readLine()?.trim()?.uppercase() ?: ""

    print("Enter years of experience: ")
    val years = readLine()?.trim()?.toIntOrNull() ?: 0

    print("Enter employee salary: ")
    val salary = readLine()?.trim()?.toDoubleOrNull() ?: 0.0

    val bonusRate = when (deptCode) {
    "IT" -> if (years > 5) 0.20 else 0.10
    "HR" -> if (years > 3) 0.15 else 0.10
    else -> 0.10
}

    val bonusAmount = salary * bonusRate

    println("Department: $deptCode")
    println("Years of Experience: $years")
    println("Bonus Rate: ${(bonusRate * 100).toInt()}%")
    println("Bonus Amount: Rs. %.2f".format(bonusAmount))
    println("Total Salary: Rs. %.2f".format(salary + bonusAmount))

    // 3. Average of Individual Digits
    println("\n--- Average of Digits ---")

    print("Enter an integer: ")
    val number = readLine()?.trim()?.toLongOrNull()

    if (number == null) {
        println("Invalid integer input.")
    } else {
        var value = if (number == Long.MIN_VALUE) {
            0L
        } else {
            kotlin.math.abs(number)
        }

        var digitSum = 0L
        var digitCount = 0

        if (number == Long.MIN_VALUE) {
            println("Number is outside the supported absolute-value range.")
        } else {
            if (value == 0L) {
                digitCount = 1
            } else {
                while (value > 0) {
                    digitSum += value % 10
                    digitCount++
                    value /= 10
                }
            }

            println("Sum of Digits: $digitSum")
            println("Number of Digits: $digitCount")
            println("Average of Digits: %.2f".format(digitSum.toDouble() / digitCount))
        }
    }
}
