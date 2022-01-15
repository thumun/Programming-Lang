package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randomTime(input int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	output := r1.Intn(input)

	// return the random number between 0-input
	return output
}

// checking if card already in collection (basically using a dictionary idea
// in slices)
// then checking if collection is complete
func checkCards(randVal int, cardCollection []int) bool {
	countSlice := 0

	if cardCollection[randVal] == 0 {
		cardCollection[randVal] = 1
	}

	for i := 0; i < len(cardCollection); i++ {
		if cardCollection[i] == 1 {
			countSlice += 1
		}
	}

	if countSlice != 34 {
		return false
	} else {
		return true
	}
}

func collect(totalCards int) int {

	// int is the number of pictures on the team
	// trying to get each person on the team

	cardCollection := make([]int, totalCards)
	randVal := randomTime(totalCards)
	allCandies := 1

	for !checkCards(randVal, cardCollection) {
		randVal = randomTime(totalCards)
		allCandies++
	}

	return allCandies

	// want to re-do function and change randomval and add one to candies

	// IDEA:
	// create a random number from 0-input value
	// then check if it exists in the card collection slice (already collected) by checking if num at array 0 or 1
	// (if 0, then didn't get card yet - if 1, then got card)
	// if 1, then add to number of candies
	// if 0, add to number of candies and change value of card collection[random] to 1
	// then change randval (call the method) and keep running through array

}

func main() {
	// testing out random function
	// for i := 0; i < 50; i++ {
	// 	a := randomTime(35)
	// 	fmt.Printf("\n%v\n", a)
	// }

	//Number of cards to collect
	n, _ := strconv.Atoi(os.Args[1])

	trials := 1000000
	sum := 0

	for i := 0; i < trials; i++ {
		sum = sum + collect(n)
		fmt.Printf("\ni: %d\n", i)

	}

	fmt.Printf("Avg. number of candies bought: %v", sum/trials)

}
