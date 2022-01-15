fun main() {
    fizzBuzz(100)
}

tailrec fun fizzBuzz(counter: Int): Unit {
    if (counter == 0) return 
    else{
        if (counter % 3 == 0 && counter % 5 == 0){
            println("FizzBuzz")
        } else if (counter % 3 == 0){
            println("Fizz")
        } else if (counter % 5 == 0){
            println("Buzz")
        } else {
            println(counter)
        }
    }

    return fizzBuzz(counter-1)
}
