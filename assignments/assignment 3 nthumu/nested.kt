/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 1 of Assignment 3.
 *    This function recursively prints the contents of 2D arrays using nested functions
 *    (using two functions)
 * Modified: Oct 29, 2021 
*/

fun main() {
    val arr = arrayOf(arrayOf(1, 2, 3), arrayOf(4, 5, 6), arrayOf(7, 8, 9))
    val arrc = arrayOf(arrayOf('c', 'd', 'e'), arrayOf('g','h','i'), arrayOf('j','l','m','n'))
    val arrs = arrayOf(arrayOf("c", "d", "e", "A"), arrayOf("g","h","i"), arrayOf("j","l","m","n", "1", "2", "3"))
    
    printArr(arr)
    printArr(arrc)
    printArr(arrs)
}    

fun <T>printArr(arr: Array<Array<T>>, defaultVal: String = "\n"): Unit {
    /* 
    * recursively cycles through the rows of given 2D array  
     */
    if (arr.size == 0){
        print(defaultVal) 
    } else{
        printMiniArr(arr[0], defaultVal) // inputs each row into this to print its contents
        val sliceArr = arr.copyOfRange(1, arr.size)
        return printArr(sliceArr, defaultVal)
    }
}

fun <T>printMiniArr(arr: Array<T>, defaultVal: String = "\n"): Unit { 
    /*
    * recursively prints the contents of given array
     */
    if (arr.size == 0){
        print(defaultVal)
    } else{
        print(arr[0])
        print(" ")
        val sliceArr = arr.copyOfRange(1, arr.size)
        return printMiniArr(sliceArr, defaultVal)
    }
}