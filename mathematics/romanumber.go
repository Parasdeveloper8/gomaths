package mathematics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type RomanNumbers struct {
	Roman   string `json:"roman"`
	English string `json:"english"`
}

// Function to get roman number
func EnglishToRoman(number string) bool {
	//Open the file
	file, err := os.Open("D:/gomaths/json/roman_english_large.json") //path to roman numbers json file
	if err != nil {
		fmt.Printf("Could not open file: %v\n", err) //could not open file
		return false
	}
	defer file.Close()
	//Read file
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Could not read file: %v\n", err)
		return false
	}
	// Parse JSON into slice of elements
	var romenNumbers []RomanNumbers
	if err := json.Unmarshal(bytes, &romenNumbers); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
		return false
	}
	for _, numeral := range romenNumbers {
		if number == numeral.English {
			fmt.Printf("Number is : %v\n and roman number is : %v\n", number, numeral.Roman)
			return true
		}
	}
	fmt.Println("could not find")
	return false
}
