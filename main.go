package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Set at build time for releases: -ldflags "-X main.version=v1.0.0"
var version = "dev"

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
	if err := json.Unmarshal(content, &compass); err != nil {
		log.Fatal("Error when parsing ", os.Args[2], ": ", err)
	}
	fmt.Println("	Looking for ", compass.Label)
	return compass
}

// Gets the index of the column with the label we are looking for
func getColumnIndex(content Compass, inputFile string) int {
	//We will use os.Open to check for the headers only, solving the issue with memory
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Error when reading ", inputFile, err)
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
		return -1
	} else {
		fmt.Println("File empty or with issues")
		return -1
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

func extractColumn(line string, targetIndex int) string {
	var (
		col      = 0     // current column counter
		startPos = 0     // byte index where the target field begins
		inField  = false // have we reached the target field yet?
	)

	// Iterate over each byte in the line
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case ',':
			// On a comma, we move to the next column
			if inField {
				// If we were in the target field, endPos is i (just before comma)
				return strings.TrimSpace(line[startPos:i])
			}
			col++
			if col == targetIndex {
				// The character after this comma is the start of our target field
				startPos = i + 1
				inField = true
			}
		case '\n', '\r':
			// End of line: if we are in the target field, slice from startPos to here
			if inField {
				return strings.TrimSpace(line[startPos:i])
			}
			// Otherwise, we've reached EOL before finding the target
			return ""
		default:
			// If we just entered the target field (col==targetIndex and inField false),
			// mark its start at the current index before any delimiter.
			if col == targetIndex && !inField {
				startPos = i
				inField = true
			}
		}
	}

	// If the loop completes, we may have the target as the last field (no trailing newline/comma)
	if inField {
		return strings.TrimSpace(line[startPos:])
	}
	// Column not found or line too short
	return ""
}

func filterCSVStream(content Compass, targetIndex int, inputFile string) error {
	// 1. The output goes next to the input: data/a.csv -> data/READY_a.csv
	outputFile := filepath.Join(filepath.Dir(inputFile), "READY_"+filepath.Base(inputFile))

	// 2. Open the input CSV for streaming read
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("opening input file %q: %w", inputFile, err)
	}
	defer inFile.Close()

	// 3. Create/truncate the output CSV
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("creating output file %q: %w", outputFile, err)
	}
	defer outFile.Close()

	// 4. Wrap in buffered I/O to keep memory usage constant
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// 5. Copy the header row verbatim
	header, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return fmt.Errorf("reading header: %w", err)
	}
	if _, err := writer.WriteString(header); err != nil {
		return fmt.Errorf("writing header: %w", err)
	}

	// 6. Stream through each subsequent row
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return fmt.Errorf("reading row: %w", err)
		}
		if line == "" {
			break // end of file
		}

		// 7. Extract only the relevant column using the provided index
		value := extractColumn(line, targetIndex)
		if value == "" {
			// skip if empty or extraction failed
			continue
		}

		// 8. Use your existing substringCheck on this single field
		if substringCheck(content.Words, value) {
			// 9. Write the full CSV row to output on match
			if _, err := writer.WriteString(line); err != nil {
				return fmt.Errorf("writing matched row: %w", err)
			}
		}
		// otherwise drop the row
		if err == io.EOF {
			break // last row had no trailing newline; it was processed above
		}
	}

	return nil
}

// Start menu
func start() {
	if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println("hankotsu", version)
		os.Exit(0)
	}

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
	columnIndex := getColumnIndex(content, os.Args[1])
	if columnIndex < 0 {
		log.Fatalf("Column %q not found in the header of %s", content.Label, os.Args[1])
	}
	fmt.Println("")
	fmt.Println("Label found at index ", columnIndex)
	if err := filterCSVStream(content, columnIndex, os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
