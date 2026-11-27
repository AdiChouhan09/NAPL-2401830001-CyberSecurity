data class Student(val name: String, val percentage: Double)

fun main() {
    val readOnlyList = listOf("Java", "Python", "Kotlin")
    println("Read-only List: $readOnlyList")
    println("Get element: ${readOnlyList[0]}")

    val mutableList = mutableListOf("Java", "Python", "Kotlin")
    mutableList.add("C++")
    mutableList.remove("Python")
    println("Mutable List: $mutableList")
    println("Contains Kotlin: ${mutableList.contains("Kotlin")}")

    val readOnlySet = setOf(10, 20, 30, 30)
    println("\nRead-only Set: $readOnlySet")

    val mutableSet = mutableSetOf(10, 20, 30)
    mutableSet.add(40)
    mutableSet.remove(20)
    println("Mutable Set: $mutableSet")
    println("Contains 30: ${mutableSet.contains(30)}")

    val readOnlyMap = mapOf(1 to "Java", 2 to "Kotlin")
    println("\nRead-only Map: $readOnlyMap")
    println("Get value: ${readOnlyMap[2]}")

    val mutableMap = mutableMapOf(1 to "Java", 2 to "Kotlin")
    mutableMap[3] = "Python"
    mutableMap.remove(1)
    println("Mutable Map: $mutableMap")
    println("Contains key 2: ${mutableMap.containsKey(2)}")

    val numbers = listOf(1, 2, 3, 4, 5, 6)
    println("\nOriginal List: $numbers")

    val mapped = numbers.map { it * 2 }
    println("Map: $mapped")

    val filtered = mapped.filter { it > 5 }
    println("Filter: $filtered")

    val reduced = filtered.reduce { sum, value -> sum + value }
    println("Reduce: $reduced")

    val students = listOf(
        Student("Aditya", 85.0),
        Student("Akansh", 72.0),
        Student("Naman", 92.0),
        Student("Kartkieya", 65.0)
    )

    val sortedStudents = students.sortedByDescending { it.percentage }

    println("\nStudents Sorted by Percentage:")
    sortedStudents.forEach {
        println("${it.name}: ${it.percentage}%")
    }

    val groupedStudents = students.groupBy {
        if (it.percentage >= 75) "Distinction"
        else "Regular"
    }

    println("\nGrouped Students:")
    groupedStudents.forEach { (group, members) ->
        println("$group: ${members.map { it.name }}")
    }
}
