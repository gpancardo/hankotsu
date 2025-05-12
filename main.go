package main

import (
	"fmt"
	"os"
)

//Struct for column label and keywords

type Compass struct{
	Label string
	Words []string
}

//Creates instance of COmpass for the current process
func loadCompass()


func main(){
	fmt.Println("")
	fmt.Println("--- HANKOTSU---")
	fmt.Println("")
	
	fmt.Println("")
	fmt.Println("Press Control + C to stop the tool at any point in time!")
	fmt.Println("( • ᴗ - )")
	fmt.Println("")

	//The program will take two arguments(input CSV file and JSON keyword list)
	//The program name counts as an argument to the os library, but it's a baseline for us
	if len(os.Args) != 3{
		fmt.Println("")
		fmt.Println("(´•︵•`)")
		fmt.Println("You should use Hankotsu as follows:")
		fmt.Println("")
		fmt.Println("	hankotsu ORIGINAL.CSV KEYWORD.JSON")
		fmt.Println("")
	}

	fmt.Println("	-⌕   Searching for ", os.Args[1])
	fmt.Println("	-⌕   Searching for ", os.Args[2])
}