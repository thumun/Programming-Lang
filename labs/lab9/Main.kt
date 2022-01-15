data class Student(val id: Int, val name: String)

fun registerStudents(students: List<String>, startId: Int = 0): List<Student>{
    if (students.size == 0) {
        return emptyList()
    } else{
        val a: Student = Student(startId, students[0])
        return [a] + registerStudents(students[1 until students.size], startId+1)
    }
}

fun main (){
    //val lst: List<Any?> = listOf(12,6,listOf(1,listOf(2)), listOf(3,4))
    //val lst: List<Any?> = listOf(listOf(12, 1), 6)
    //println(flatten(lst))

    val students = listOf("Alice", "Bob")
    registerStudents(students)

    registerStudents(students, startId = 10)
}

//fun flatten(lst: List<Any?>): List<Any?>{
//    if (lst.size == 0){
//        return emptyList()
//    } else {
//        if (lst[0] !is List<Any?>) {
//            return listOf(lst[0]) + flatten(lst.slice(1..lst.size - 1))
//        } else {
//            val miniLst: List<Any?> = lst[0] as List<Any?>
//            return flatten(miniLst) + flatten(lst.slice(1..lst.size - 1))
//        }
//    }
//}

