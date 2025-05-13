package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//keywords.json sample:
//{
// 	"label":"sample_label",
// 	"words":["first_word","second_word"]
//}

// Struct for column label and keywords
type Compass struct {
	Label string
	Words []string
}

// Reads JSON information and loads it to compass
func loadCompass() (currentCompass Compass) {
	content, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal("Error when reading ", os.Args[2], err)
	}
	var compass Compass
	//json.Unmarshal writes the JSON content straight into the struct instance
	json.Unmarshal(content, &compass)
	fmt.Println("	Looking for ", compass.Label)
	return compass
}

// Gets the index of the column with the label we are looking for
func getColumnIndex(content Compass) int {
	//We will use os.Open to check for the headers only, solving the issue with memory
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error when reading ", os.Args[1], err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		//Gets header row
		headersLine := scanner.Text()
		fmt.Println("	File headers: ", headersLine)
		//The header row string becomed a list with each element divided by a comma
		listHeaders := strings.Split(headersLine, ",") //Checks for match
		for i := 0; i <= (len(listHeaders) - 1); i++ {
			if listHeaders[i] == content.Label {
				return i
			}
		}
		return 100
	} else {
		fmt.Println("File empty or with issues")
		return 101
	}
}

// Check if a string has a substring from the list we are interested in
func substringCheck(keywordList []string, originalValue string) bool {
	for j := 0; j <= (len(keywordList) - 1); j++ {
		if strings.Contains(originalValue, keywordList[j]) {
			return true
		}
	}
	return false
}

// Start menu
func start() {
	fmt.Println("")
	fmt.Println("--- HANKOTSU---")
	fmt.Println("")

	fmt.Println("")
	fmt.Println("Press Control + C to stop the tool at any point in time!")
	fmt.Println("( • ᴗ - )")
	fmt.Println("")

	//The program will take two arguments(input CSV file and JSON keyword list)
	//The program name counts as an argument to the os library, but it's a baseline for us
	if len(os.Args) != 3 {
		fmt.Println("")
		fmt.Println("You should use Hankotsu as follows:")
		fmt.Println("(´•︵•`)")
		fmt.Println("")
		fmt.Println("	hankotsu ORIGINAL.CSV KEYWORD.JSON")
		fmt.Println("")
		os.Exit(0)
	}

	fmt.Println("	-⌕   Searching for ", os.Args[1])
	fmt.Println("	-⌕   Searching for ", os.Args[2])
}

func main() {
	start()
	content := loadCompass()
	columnIndex := getColumnIndex(content)
	fmt.Println("")
	fmt.Println("Label found at index ", columnIndex)
}
