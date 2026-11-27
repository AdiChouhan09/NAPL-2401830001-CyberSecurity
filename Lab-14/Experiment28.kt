
fun main() {
    val nonNullable: String = "Kotlin"
    var nullable: String? = "Programming"

    println("Non-nullable String: $nonNullable")
    println("Nullable String: $nullable")

    nullable = null
    println("Nullable String after assigning null: $nullable")

    val safeLength = nullable?.length
    println("\nSafe Call Result: $safeLength")

    val message: String? = "Hello Kotlin"
    println("Length of Message: ${message?.length}")

    val defaultValue = nullable ?: "Default Value"
    println("\nElvis Operator Result: $defaultValue")

    val defaultLength = nullable?.length ?: 0
    println("Default Length: $defaultLength")

    val nullValue: String? = null

    try {
        println(nullValue!!.length)
    } catch (e: NullPointerException) {
        println("\nNullPointerException caught successfully.")
    }

    println("Program completed successfully.")
}
