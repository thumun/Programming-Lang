/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 2 of Assignment 1.
 *    This program is creating slices of the data from part 1 and
 *	  manipulating splices (creating subsplices).
 * Modified: Sep 19, 2021
 */

package main

import (
	"fmt"
)

/**
 * Type definition for arrival/departure times of airplanes
**/
type ArrDepTimes struct {
	year    int
	month   int
	day     int
	deptime int
	arrtime int
}

/**
* String format function for ArrDepTimes
**/
func (adt ArrDepTimes) String() string {
	output := fmt.Sprintf("\tyear:%d month:%d day:%d deptime:%d arrtime:%d \n",
		adt.year, adt.month, adt.day, adt.deptime, adt.arrtime)
	return output
}

/**
 * Type definition for airplane information
**/
type Aircraft struct {
	carrier  string
	fightnum int
}

/**
 * String format function for Aircraft
**/
func (ac Aircraft) String() string {
	output := fmt.Sprintf("\tcarrier:%s flightnum:%d \n", ac.carrier, ac.fightnum)
	return output

}

/**
 * Type definition for all available flight information
**/
type AircraftActivities struct {
	arrDepTimes ArrDepTimes
	aircraft    Aircraft
	arrdelay    int
	depdelay    int
	origin      string
	destination string
	distance    int
	cancelled   bool
}

/**
 * String format function for AircraftActivities
**/
func (aa AircraftActivities) String() string {
	output := fmt.Sprintf("\narrDepTimes:\n%vaircraft:\n%varrdelay:%v depdelay:%v origin:%v destination:%v distance:%v cancelled:%v \n",
		aa.arrDepTimes, aa.aircraft, aa.arrdelay, aa.depdelay, aa.origin,
		aa.destination, aa.distance, aa.cancelled)

	return output
}

func main() {
	// creating objects for Aircraft Activities
	a := AircraftActivities{ArrDepTimes{2008, 1, 3, 1734, 1941},
		Aircraft{"WN", 23}, 36, 44, "JAX", "PHL", 742, false}
	b := AircraftActivities{ArrDepTimes{2008, 1, 3, 712, 926},
		Aircraft{"WN", 1232}, 11, 12, "JAX", "PHL", 742, false}
	c := AircraftActivities{ArrDepTimes{2008, 1, 3, 1127, 1856},
		Aircraft{"WN", 1285}, 21, 42, "LAS", "PHL", 2176, true}

	// slice containing objects of AircraftActivites
	s := make([]AircraftActivities, 3)
	s[0] = a
	s[1] = b
	s[2] = c

	// adding the rest of the data to the slice
	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 12, 1834, 2149},
		Aircraft{"DL", 1678}, 86, 69, "PHL", "SLC", 1926, true})

	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 12, 944, 1555},
		Aircraft{"DL", 1776}, -10, -1, "SLC", "PHL", 1926, true})

	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 13, 1221, 1413},
		Aircraft{"DL", 692}, -9, -4, "ATL", "PHL", 665, true})

	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 5, 1523, 1804},
		Aircraft{"AA", 1768}, -21, 3, "ORD", "PHL", 678, true})

	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 6, 1628, 1933},
		Aircraft{"AA", 1768}, -7, -2, "ORD", "PHL", 678, true})

	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 7, 1520, 1804},
		Aircraft{"AA", 1768}, -21, 0, "ORD", "PHL", 678, true})

	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 22, 545, 839},
		Aircraft{"CO", 1677}, 19, 10, "PHL", "IAH", 1324, true})

	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 22, 1920, 2314},
		Aircraft{"CO", 1776}, -16, 0, "IAH", "PHL", 1324, true})

	s = append(s, AircraftActivities{ArrDepTimes{2008, 12, 22, 723, 952},
		Aircraft{"CO", 1777}, 3, 18, "PHL", "IAH", 1324, true})

	// first five items in slice
	firstFive := s[0:5]
	fmt.Printf("\nfristfive:\n%v\n", firstFive)
	fmt.Printf("fristfive length: %v\n", len(firstFive))

	// last 5 items
	lastFive := s[len(s)-5:]
	fmt.Printf("\nlastfive:\n%v\n", lastFive)
	fmt.Printf("lastFive length: %v\n", len(lastFive))

	// first three items and last two items
	threeAndTwo := s[0:3]
	threeAndTwo = append(threeAndTwo, s[len(s)-2:]...)

	fmt.Printf("\nthreeAndTwo:\n%v\n", threeAndTwo)
	fmt.Printf("length threeAndTwo: %v\n", len(threeAndTwo))
}
