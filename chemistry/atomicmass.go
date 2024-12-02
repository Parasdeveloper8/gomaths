package chemistry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Element struct {
	Name        string  `json:"name"`
	Atomic_Mass float64 `json:"atomic_mass"`
	Symbol      string  `json:"symbol"`
}

// GetAtomicMass retrieves the atomic mass of an element by name or symbol.
func GetAtomicMass(name string) (float64, bool) {

	// Open the elements file
	file, err := os.Open("D:/gomaths/json/elements.json") // Path to elements.json
	if err != nil {
		fmt.Printf("Could not open file: %v\n", err) //Could not open file
		return 0, false
	}
	defer file.Close()

	// Read file contents
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Could not read file: %v\n", err) //Could not read file contents
		return 0, false
	}

	// Parse JSON into slice of elements
	var elements []Element
	if err := json.Unmarshal(bytes, &elements); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err) //parse to json failed
		return 0, false
	}

	// Search for the element by name or symbol
	for _, element := range elements {
		if (len(name) == 2 || len(name) < 2 && name == element.Symbol) || (len(name) > 2 && name == element.Name) {
			fmt.Printf("Element: %v\n Mass: %v\nSymbol: %v\n", element.Name, element.Atomic_Mass, element.Symbol)
			return element.Atomic_Mass, true
		}

	}

	// If element is not found
	fmt.Println("Element not found.")
	return 0, false
}
