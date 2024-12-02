package mathematics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type FormulaStore struct {
	Name        string `json:"name"`
	Formula     string `json:"formula"`
	Description string `json:"description"`
	Topic       string `json:"topic"`
}

// Get formulas of related given topic
func GetFormula(topic string) (bool, float64) {
	//open formulas file
	file, err := os.Open("D:/gomaths/json/formulas.json") //path to file
	if err != nil {
		log.Fatal("Could not open file", err) //If file could not be opened
		return false, 0
	}
	defer file.Close()

	//Read file contents
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Could not read file: %v\n", err) //could not read file contents
		return false, 0
	}

	// Parse JSON into slice of elements
	var topicName []FormulaStore
	if err := json.Unmarshal(bytes, &topicName); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err) //Failed to parse json
		return false, 0
	}
	for _, formula := range topicName {
		if topic == formula.Name || topic == formula.Topic {
			fmt.Printf("name : %v\nformula : %v\n description : %v\n topic : %v\n", formula.Name, formula.Formula, formula.Description, formula.Topic)
			return true, 1
		}
	}
	fmt.Println("Could not find formula : use First Letter capital or formula is not available")
	return false, 0
}
