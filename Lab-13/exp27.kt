class Student(
    var rollNo: Int,
    var name: String,
    var course: String,
    var percentage: Double
) {

    fun updateData() {
        print("Enter Roll Number: ")
        rollNo = readLine()!!.toInt()

        print("Enter Name: ")
        name = readLine()!!

        print("Enter Course: ")
        course = readLine()!!

        print("Enter Percentage: ")
        percentage = readLine()!!.toDouble()
    }

    fun displayDetails() {
        println("\nStudent Details")
        println("Roll Number: $rollNo")
        println("Name: $name")
        println("Course: $course")
        println("Percentage: $percentage%")
    }
}

fun main() {
    val student1 = Student(101, "Aditya", "B.Sc. Cyber Security", 85.5)
    val student2 = Student(102, "Rahul", "B.Sc. Computer Science", 90.0)

    student1.updateData()
    student1.displayDetails()

    student2.updateData()
    student2.displayDetails()
}