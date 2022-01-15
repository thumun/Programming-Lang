fun main(args: Array<String>){
    convert(args[0], args[1])

}

fun convert(strInput: String, unit: String): Unit{
    if (unit == "k"){
        val intInput: Int = strInput.toInt()
        val output: Double = intInput / 1.62137119
        println("$output miles") 
    } else if (unit == "m"){
        val intInput: Int = strInput.toInt()
        val output: Double = intInput * 1.62137119
        println("$output kilometers") 
    }
}