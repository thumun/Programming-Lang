/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 1 of Assignment 3.
 *    This function recursively prints hello world with the intention of 
 *    getting stack overflow. 
 *    This recursion function does not use tail recursion. 
 * Modified: Oct 29, 2021 
*/

fun main(args: Array<String>){
    val count: Int = args[0].toInt() // converts input into an integer 

    printHW(count)
}

fun printHW(va: Int): String{
    /**
	* recursively prints hello world based on input
	**/
    if (va == 0){
        return ""
    } else {
        print("Hello World\n")
        return printHW(va-1) 
    }
}