/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 1 of Assignment 1.
 *    This program is organizing airplane data into structs.
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

	// comparing the values of the objects
	fmt.Printf("a:%v, a==b:%t\n", a, a == b)
	fmt.Printf("b:%v, b==c:%t\n", b, b == c)
	fmt.Printf("c:%v, c==a:%t\n", c, c == a)

	d := a // copy of a
	fmt.Printf("d:%v, d==a:%t\n", d, d == a)

	d.destination = "TNT" // hopefully trenton airport initials

	// testing if a and d still equal
	fmt.Printf("d:%v, d==a:%t\n", d, d == a)

	// The original and the copy are not equal.
	// They would be equal if made in Java. Because in Java,
	// d and a would be pointing to the same place in memory.
}
