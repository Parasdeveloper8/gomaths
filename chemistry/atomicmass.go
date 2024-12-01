package chemistry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Element struct {
	Name string  `json:"name"`
	Mass float64 `json:"mass"`
}

func GetAtomicMass(name string) (float64, bool) {
	file, err := os.Open("D:/gomaths/json/elements.json")
	if err != nil {
		fmt.Printf("Couldnot open file %v", err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Could not read file %v", err)
	}
	var elements []Element
	if err := json.Unmarshal(bytes, &elements); err != nil {
		log.Fatalf("Failed to parse JSON %v", err)
	}

	for _, element := range elements {
		if element.Name == name {
			fmt.Printf("Element %v\n : Mass %v \n", name, element.Mass)
			return element.Mass, true
		}
		//fmt.Printf("Element name : %v\n Element mass :%v\n", element.Name, element.Mass)
	}
	return 0, false
}
