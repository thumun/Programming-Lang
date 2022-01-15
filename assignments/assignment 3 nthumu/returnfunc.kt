/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 1 of Assignment 3.
 *    This function recursively calculates the exponent of a base using a 
 *    function and its closure. 
 * Modified: Oct 29, 2021 
*/

fun main() {
    val pow4 = nthpower(4)
    println(pow4(2))
    println(pow4(3))
    println(pow4(4))
    println(pow4(5))

    println(nthpower(5)(2))
    println(nthpower(8)(2))
}   

typealias Power = (Int)->Int
typealias Exp = (Int, Int)->Int

fun nthpower(exponent: Int): Power{
    /* 
    * calculating the power of a number 
    */

    // gotta use vars here to do closures
    lateinit var power: Power 
    lateinit var expfunc: Exp

    power = fun (base: Int): Int{
        expfunc = fun(b: Int, e: Int): Int{
            if (e == 0){
                return 1
            }else{
                return b * expfunc(b, e-1)
            }
        }
        return expfunc(base, exponent)
    }
    return power  
}

