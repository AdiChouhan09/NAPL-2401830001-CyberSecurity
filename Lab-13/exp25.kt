fun main() {

    // 1. Hello Kotlin
    println("Hello Kotlin!!!")

    // 2. Employee Bonus
    print("Enter department code: ")
    val dept = readLine()!!.uppercase()

    print("Enter years of experience: ")
    val years = readLine()!!.toInt()

    print("Enter salary: ")
    val salary = readLine()!!.toDouble()

    val bonusPercent = when {
        dept == "IT" && years > 5 -> 20
        dept == "HR" && years > 3 -> 15
        dept == "SALES" && years > 2 -> 12
        else -> 10
    }

    val bonus = salary * bonusPercent / 100

    println("Bonus Percentage: $bonusPercent%")
    println("Bonus Amount: $bonus")

    // 3. Average of Individual Digits
    print("Enter an integer: ")
    val num = readLine()!!.toInt()

    var n = kotlin.math.abs(num)
    var sum = 0
    var count = 0

    if (n == 0) {
        count = 1
    }

    while (n > 0) {
        val digit = n % 10
        sum += digit
        count++
        n /= 10
    }

    val average = sum.toDouble() / count

    println("Average of digits: $average")
}