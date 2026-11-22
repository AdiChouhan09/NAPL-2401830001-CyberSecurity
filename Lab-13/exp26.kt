fun greet(name: String, message: String = "Hello", city: String = "Delhi") {
    println("$message, $name from $city")
}

fun add(a: Int, b: Int): Int {
    return a + b
}

fun printNumbers(vararg numbers: Int) {
    for (num in numbers) {
        println(num)
    }
}

fun applyOperation(numbers: List<Int>, operation: (Int) -> Int) {
    for (num in numbers) {
        println(operation(num))
    }
}

fun main() {
    greet("Aditya")
    greet("Aditya", "Good morning", "Delhi")

    println("Sum: ${add(b = 20, a = 10)}")

    printNumbers(10, 20, 30, 40)

    val numbers = listOf(1, 2, 3, 4, 5)
    applyOperation(numbers) { it * it }
}