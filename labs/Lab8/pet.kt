open class Pet(val Id: String, val Name:String){
    //val ID: String = this.ID 
    //val name: Double = Name

    open fun sound(): String {
        return "pet sound"
    }

    override fun toString() = String.format("Id: %s \n Name: %s", Id, Name)
}

class Cat(Id: String, val Breed: String, Name: String, val hairLength: Double): Pet(Id, Name){
    //val breed: String = Breed 
    //val hairLength: Double = this.hairLength

    override fun sound(): String{
        return "meow"
    }

    override fun toString() = String.format("ID: %s \n Breed: %s \n Name: %s \n Hair Length: %f", Id, Breed, Name, hairLength)
}

open class Dog(Id: String, val Group: String, val Breed: String, Name: String, val hairLength: Double, val doubleCoat: Boolean) : Pet(Id, Name){
    // val group: String = Group 
    // val breed: String = Breed 
    // val hairlength: Double = this.hairLength
    // val doubleCoat: Boolean = this.doubleCoat 
    
    override fun sound(): String{
        return "woof"
    }

    override fun toString() = String.format("ID: %s \n Group: %s \n Breed: %s \n Name: %s \n Hair length: %f \n Double coat: %b", Id, Group, Breed, Name, hairLength, doubleCoat)
}

class WorkingDog(Id: String, Group: String, Breed: String, Name: String, hairLength: Double, doubleCoat: Boolean, val typeOfWork: String) : Dog(Id, Group, Breed, Name, hairLength, doubleCoat){
    // val typeOfWork: String = this.typeOfWork
    
    override fun sound(): String{
        return "italicized woof"
    }

    override fun toString() = String.format("ID: %s \n Group: %s \n Breed: %s \n Name: %s \n Hair length: %f \n DoubleCoat: %b \n Type of Work: %s", Id, Group, Breed, Name, hairLength, doubleCoat, typeOfWork)

}

fun main(){
    val liz: Pet = Pet("1310", "Liz")
    println(liz)
    println(liz.sound())

    val bob: Cat = Cat("805", "orange?", "BOB", 12.0)
    println(bob)
    println(bob.sound())

    val riot: Dog = Dog("801", "large", "Australian Shepherd", "Riot", 18.0, true)
    println(riot)
    println(riot.sound())

    val shiro: Dog = WorkingDog("423", "small", "Japanese Spitz", "Shiro", 12.0, true, "show dog")
    println(shiro)
    println(shiro.sound())

    val newShiro: Dog = shiro 
    println(newShiro)

    // prints out the working dog instead of the dog contents. This shows dynamic dispatch. If it was static dispath, would have 
    // printed the dog toString. 
}