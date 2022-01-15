/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 1 of Assignment 3.
 *    This filters elements of a list 
 * Modified: Oct 29, 2021 
*/

fun main() {
    val ll = listOf(1,2,3,4,5,6,7,8,9)

    println(ll.myFilter(isOdd))
    println(ll.myFilter(isEven))
    println(ll.myFilter(isDivisibleByTwoThree))
}

val isOdd = fun (n:Int) : Boolean {
    /* 
    * checks if the given element is odd
    */
    return (n%2!=0)
}

val isEven = fun(n: Int) : Boolean {
    /* 
    * checks if the given element is even
    */
    return (n%2==0)
}

val isDivisibleByTwoThree = fun(n:Int) : Boolean {
    /* 
    * checks if the given element is divisible by 2 and 3
    */
    return (n%2==0 && n%3==0)
}

typealias Select = (List<Int>)->List<Int>

fun List<Int>.myFilter(funcc: (Int) -> Boolean) : List<Int> {
    /* 
    * applying a function to each element of list and creating a list 
    * containing the elements that fulfill the function
    */

    val filteredList: MutableList<Int> = mutableListOf<Int>()

    lateinit var select: Select

    select = fun (inputLst:List<Int>): List<Int>{
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