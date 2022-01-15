/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 1 of Assignment 3.
 *    This filters elements of a list (of generic type)
 * Modified: Oct 29, 2021 
*/

fun main() {
    val ll = listOf(1,2,3,4,5,6,7,8,9)
    val doubleInputs = listOf(0.5, 0.3, 7.9, 10.3, 5.3, 5.6)
    val strInputs = listOf("hi", "Bye", "Neha", "Thumu", "Pokemon", "a")

    println(ll.myFilterGeneric(isOdd))
    println(doubleInputs.myFilterGeneric(isGreaterThanFive))
    println(strInputs.myFilterGeneric(isShortString))
}

val isOdd = fun (n:Int) : Boolean {
    /* 
    * checks if the given element is odd
    */
    return (n%2!=0)
}

val isGreaterThanFive = fun(d: Double) : Boolean {
    /* 
    * checks if the given element is even
    */
    return (d > 5)
}

val isShortString = fun(s:String) : Boolean {
    /* 
    * checks if the given element is divisible by 2 and 3
    */
    return (s.length <= 2)
}

typealias Select<Q> = (List<Q>)->List<Q>

fun <Q> List<Q>.myFilterGeneric(funcc: (Q) -> Boolean) : List<Q> {
    /* 
    * applying a function to each element of list and creating a list 
    * containing the elements that fulfill the function
    */

    val filteredList: MutableList<Q> = mutableListOf<Q>()

    lateinit var select: Select<Q>

    select = fun (inputLst:List<Q>): List<Q>{
        if (inputLst.size == 0){
            return filteredList
        }
        else{
            if (funcc(inputLst[0])){
                filteredList.add(inputLst[0])
            } 
            val num: Int = inputLst.size - 1
            return select(inputLst.slice(1..num))
        }
    }
    return select(this)
}