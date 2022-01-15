/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Assignment 4.
 *    Has the contents of the assignment.
 *    Reading data from a file and organizing the data into objects of a class.
 *    Then filtering the data such that the correct output is printed based on the command line arguments.
 * Modified: Nov 15, 2021
*/

import java.io.BufferedReader
import java.io.InputStreamReader
import java.net.URL

fun main(args: Array<String>) {
    val years = (2000..2009)
    // string intarporltation to get and read files
    val filesContents = years.map{"https://cs.brynmawr.edu/cs245/Data/names$it.csv"}.map{readURL(it)}
    val parsedData = filesContents.map{readList(it)}

    val allBoys = parsedData.map{it.first}
    val allGirls = parsedData.map{it.second}

    args.forEach{ input:String ->
        println(input)

        // merges all the (boy) results into one list and groups together the same names
        val boyResults = allBoys.map { it.searchBy(input) }.flatten().groupBy { it.name }

        // printing results based on output from combineSearchResults
        boyResults.forEach{
            println(combineSearchResults(it.value))
        }
        val girlResults = allGirls.map { it.searchBy(input) }.flatten().groupBy { it.name }

        girlResults.forEach{
            println(combineSearchResults(it.value))
        }
        println()
    }
}

/*
* reads the data from a file
 */
val readURL = { url:String ->
    val r = mutableListOf<String>();
    try {
        val oracle = URL(url);
        val br = BufferedReader(InputStreamReader(oracle.openStream()));
        loop@ while (true) {
            val line = br.readLine();
            if (line==null) {
                break@loop
            }
            r.add(line);
        }
    } catch (ee: Exception) {
        println("Problem ${ee}");
    }
    r
}

/*
* if there are appearances of the same name in multiple years, will combine the data
 */
fun combineSearchResults(searchResultLists: List<SearchResult>): SearchResult {
    val freqSum = searchResultLists.sumOf{it.freq}
    val totalFreqSum = searchResultLists.sumOf{it.totalFreq}
    val alphPosition = searchResultLists.minOf{it.alphPosition}

    return SearchResult(alphPosition, searchResultLists[0].gender, searchResultLists[0].name, searchResultLists.size,
        freqSum, totalFreqSum)
}

/* holds the contents of the list made up of NameHolders from each year for a gender
 */
class NameHolderCollection(private val gender: String, val names: List<NameHolder>){

    // sorting the list alphabetically
    private val sortedCollection = names.sortedBy{ it.name }
        .mapIndexed{ index, item -> NameHolder(index+1, item.name, item.freq) }

    // sum of all the frequencies in the collection
    private val totalFreq = sortedCollection.sumOf { it.freq }

    // generating search results based on the search term
    fun searchBy(searchTerm: String): List<SearchResult> {
        return sortedCollection.filter{it.name.startsWith(searchTerm, ignoreCase = true)}
            .map {SearchResult(it.row, gender, it.name, 1, it.freq, totalFreq)}
    }
}

/*
* contains the relevant information for each item (name) in order to set up the correct print format
 */
data class SearchResult(val alphPosition: Int, val gender: String, val name: String, val yearsUsed: Int, val freq: Int,
                        val totalFreq: Int){

     private val percentUsed = (freq.toFloat()/totalFreq.toFloat())*100

    override fun toString(): String {
        return "%-4s\t%-10s\t\t%2d\t%5d\t%.4f\n\talphabetic position %-5d"
            .format(gender, name, yearsUsed, freq, percentUsed, alphPosition)
    }
}

/*
* class that has objects for each name and the name's relevant information
 */
data class NameHolder(val row: Int, val name: String, val freq: Int){
}

/*
* reading the data into a workable format (making them objects of NameHolder and creating a pair of
* NameHolderCollections -- a collection of boy names and a collections of girl names for a year)
 */
fun readList(nameList: MutableList<String>): Pair<NameHolderCollection, NameHolderCollection> {
    val allNames = nameList.map {
        mapNames(it)
    }

    val boys = NameHolderCollection("Boy", allNames.map{it.first})
    val girls = NameHolderCollection("Girl", allNames.map{it.second})

    return Pair(boys, girls)
}

/*
* Helper function to split one line from file into pair of boy and girl objects
 */
fun mapNames(line: String): Pair<NameHolder, NameHolder>{
    val listTest = line.split(",")
    val boy = NameHolder(listTest[0].toInt(), listTest[1], listTest[2].toInt())
    val girl = NameHolder(listTest[0].toInt(), listTest[3], listTest[4].toInt())

    return Pair(boy, girl)
}