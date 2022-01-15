package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

// holding the names & their data from reading the files
type nameHolder struct {
	row  int
	name string
	freq int
}

// methods for sorting by name on the nameHolders
type byName []nameHolder

func (n byName) Len() int           { return len(n) }
func (n byName) Less(i, j int) bool { return n[i].name < n[j].name }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

// holding nameHolders by years (files) and gender
type nameHolderCollection struct {
	gender string
	names  [][]nameHolder
}

// relevant fields for output
type searchResult struct {
	alphPos   int
	gender    string
	name      string
	yearsUsed int
	freq      int
	totalFreq int
}

func (n *nameHolderCollection) add(names []nameHolder, gender string) {
	/*
	*  adding names (as type nameHolder) to nameHolderCollection
	*  also sorting the names based on alphabetical order
	 */
	sort.Sort(byName(names))
	for index, _ := range names {
		// changing row to reflect the name's popularity
		names[index].row = index + 1 // +1 b/c index starts from 0
	}

	n.names = append(n.names, names)
	n.gender = gender
}

func (n *nameHolderCollection) searchBy(searchTerm string, gender string) []searchResult {
	/*
	* searching for the searchterm (a name) in the names from the different files
	 */

	var output []searchResult

	for i, val := range n.names {
		// calculating the total frequency for each i (file)
		var totalFreq int
		for tf := 0; tf < len(val); tf++ {
			totalFreq += n.names[i][tf].freq
		}

		for j := 0; j < len(val); j++ {
			// checking if strings same (case insenstive) and looking for names that has a prefix
			if strings.EqualFold(n.names[i][j].name, searchTerm) || strings.HasPrefix(strings.ToLower(n.names[i][j].name), strings.ToLower(searchTerm)) {
				output = append(output, searchResult{n.names[i][j].row, gender, n.names[i][j].name, 1, n.names[i][j].freq, totalFreq})
			}
		}
	}

	return output
}

func combineDistinct(results []searchResult) map[string]searchResult {
	/*
	*  combining search results of the same name into one value (adding their freqs, yearsUsed, totalfreqs)
	 */

	// making a map with keys as names and values as searchResult [name info in format that will be used in output]
	m := make(map[string]searchResult)
	for _, r := range results {
		val, err := m[r.name]

		// if name not already in map, add it
		if !err {
			m[r.name] = r
		} else {
			var alphPos int

			// taking the highest alpabetical value
			if val.alphPos > r.alphPos {
				alphPos = r.alphPos
			} else {
				alphPos = val.alphPos
			}

			m[r.name] = searchResult{alphPos, r.gender, r.name, r.yearsUsed + val.yearsUsed, r.freq + val.freq, r.totalFreq + val.totalFreq}
		}
	}
	return m
}

func formatResult(s searchResult) string {
	/*
	* calculating the total frequency and formatting the output string with proper values & appearance
	 */

	// calculating the total frequency percentage
	percentUsed := (float64(s.freq) / float64(s.totalFreq)) * 100

	return fmt.Sprintf("%v\t%v\t\t\t%v\t%v\t%.4f\n\t\talphabetic position %v", s.gender, s.name, s.yearsUsed, s.freq, percentUsed, s.alphPos)
}

func readHurricane(url string, maxLines int) ([]nameHolder, []nameHolder) {
	/*
	* read function given by Prof Towell
	* [I took out the print statements but otherwise mostly same]
	 */

	// getting the files from the site
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http problem", err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Http problem2", err)
	}
	// all data in one string
	str := string(body)
	r := csv.NewReader(strings.NewReader(str))
	r.FieldsPerRecord = -1

	var girlNames []nameHolder
	var boyNames []nameHolder

	// reading files line by line
	// each line is a name so making it into a nameHolder struct
	// then adding the nameHolders to a slice based on gender
	for i := 0; i < maxLines || maxLines <= 0; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		// nameholder obj
		rowNum, _ := strconv.Atoi(record[0])
		boyFreq, _ := strconv.Atoi(record[2])
		boy := nameHolder{rowNum, record[1], boyFreq}

		boyNames = append(boyNames, boy)

		girlFreq, _ := strconv.Atoi(record[4])
		girl := nameHolder{rowNum, record[3], girlFreq}

		girlNames = append(girlNames, girl)
	}
	return boyNames, girlNames
}

func main() {
	var allBoys nameHolderCollection
	var allGirls nameHolderCollection

	// getting the name inputs for the search
	nameInputs := os.Args[1:]

	// reading the files and organizing the names based on gender
	for i := 2000; i <= 2009; i++ {
		boyNames, girlNames := readHurricane(fmt.Sprintf("https://cs.brynmawr.edu/cs245/Data/names%d.csv", i), 1000)

		allBoys.add(boyNames, "Boy")
		allGirls.add(girlNames, "Girl")
	}

	// searching each input
	for _, val := range nameInputs {
		boy := allBoys.searchBy(val, allBoys.gender)
		girl := allGirls.searchBy(val, allGirls.gender)

		// combining results of the same name (from diff. files)
		outputBoy := combineDistinct(boy)
		outputGirl := combineDistinct(girl)

		fmt.Println(val)
		// printing the results in the desired format
		for _, valBoy := range outputBoy {
			fmt.Println(formatResult(valBoy))
		}

		for _, valGirl := range outputGirl {
			fmt.Println(formatResult(valGirl))
		}

		fmt.Println()
	}

}
