package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type dateofBirth struct {
	Year  int
	Month string
	Day   int
}

type Student struct {
	FirstName          string
	LastName           string
	Id                 int
	Age                int
	DOB                dateofBirth
	FavoriteRealNumber int
}

type jm map[string]interface{}
type js []jm

// type studentSlice []Student

func main() {
	jsonFile, err := os.Open("/Users/nehathumu/Desktop/Ursae Majoris/Haverford/2021-2022/cs245/labs/lab 6/users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	// var res js
	// json.Unmarshal([]byte(byteValue), &res)
	// fmt.Printf("%v", res)

	var students []Student
	json.Unmarshal([]byte(byteValue), &students)
	//fmt.Printf("%v", students)

	data, _ := json.Marshal(students)
	fmt.Println(string(data))

}
