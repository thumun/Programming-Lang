/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 1 of Assignment 3.
 *    This function recursively prints the contents of 2D arrays using nested functions
 *    (using one functions)
 * Modified: Oct 29, 2021 
*/

fun main() {
    val arr = arrayOf(arrayOf(1, 2, 3), arrayOf(4, 5, 6), arrayOf(7, 8, 9))
    val arrc = arrayOf(arrayOf('c', 'd', 'e'), arrayOf('g','h','i'), arrayOf('j','l','m','n'))
    val arrs = arrayOf(arrayOf("c", "d", "e", "A"), arrayOf("g","h","i"), arrayOf("j","l","m","n", "1", "2", "3"))
    
    printArrV2(arr)
    printArrV2(arrc)
    printArrV2(arrs)
}    

fun <T>printArrV2(arr: Array<Array<T>>, defaultVal: String = "\n"): Unit {
    /* 
    * recursively cycles through the rows of given 2D array  
     */
    if (arr.size == 0){
        print(defaultVal) 
    } else{
        // takes the first element of the first row of 2D array then changes the row
        // to remove the first element 
        print(arr[0][0])
        print(" ")
        arr[0] = arr[0].copyOfRange(1, arr[0].size)

        // if the row is empty, then eliminates the row from the 2D array
        if (arr[0].size == 0){
            val sliceArr = arr.copyOfRange(1, arr.size)
            print(defaultVal)
            return printArrV2(sliceArr, defaultVal)
        } else{
            return printArrV2(arr, defaultVal)
        }
    }
}
