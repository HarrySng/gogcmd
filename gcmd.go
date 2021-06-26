package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type term struct {
	id   string
	path string
}

// Taken from the GCMD static page here: https://gcmd.earthdata.nasa.gov/static/kms/
var concepts = []string{
	"sciencekeywords", "providers", "platforms", "mimetype",
	"locations", "chronounits", "projects", "rucontenttype",
	"verticalresolutionrange", "horizontalresolutionrange",
	"temporalresolutionrange", "granuledataformat", // Mandatory comma at end if slice split in multiple lines
}

func main() {

	// Create txt files
	for _, concept := range concepts {
		downloadCSV(concept)
		data := createTXT(concept) // data is a slice of strings
		terms := parseData(data)
		writeTerms(terms, concept)
	}
}

func parseData(data []string) []term {

	terms := []term{}

	for i, d := range data {
		if i < 2 {
			continue
		}
		term := curate(d)
		terms = append(terms, term)
	}
	return terms
}

func curate(d string) term {
	d = strings.Replace(d, "\"", "", -1) // Remove double quotes in string
	fields := strings.Split(d, ",")      // Split by comma
	fields = removeBlanks(fields)        // Use helper function to remove blank strings
	if len(fields) < 3 {
		return term{id: fields[1], path: fields[0]}
	}
	id := fields[len(fields)-1]                          // ID is the last field
	fields = fields[0 : len(fields)-1]                   // Drop the ID string
	path := strings.Join(fields[0:len(fields)-1], " > ") // Join all strings except last
	path = path + " > " + fields[len(fields)-1]          // Now join the last. This is to avoid extra '>' at the end
	term := term{id: id, path: path}
	return term
}

func removeBlanks(fields []string) []string {
	var flds []string
	for _, s := range fields {
		if s != "" {
			flds = append(flds, s)
		}
	}
	return flds
}

func writeTerms(terms []term, concept string) {

	filepath := "./terms/" + concept + ".txt"

	// Create a file in none exists and append to file when writing
	if _, err := os.Stat(filepath); err == nil {
		fmt.Println("File for " + concept + " already exists")
		return
	}

	f, err := os.Create(filepath)
	handleError(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	for _, term := range terms {
		w.WriteString(term.id + "\n")
		w.WriteString(term.path + "\n\n")
	}
	w.Flush()
}

func createTXT(concept string) []string {

	/*
		Creating txt files even though data is csv
		because the csv format is irregular in terms of
		number of columns in each row.
		It is easier to read each row as a single string
		and then parse it.
	*/

	f, err := os.Open("./files/" + concept + ".txt")
	handleError(err)
	defer f.Close()

	scn := bufio.NewScanner(f)
	scn.Split(bufio.ScanLines) // Split lines by new line
	var data []string

	for scn.Scan() {
		data = append(data, scn.Text()) // Append each line to slice
	}

	return data
}

func downloadCSV(concept string) {

	url := "https://gcmd.earthdata.nasa.gov/kms/concepts/concept_scheme/" + concept + "/?format=csv"
	filepath := "./files/" + concept + ".txt"

	// Do not download if the file already exists
	if _, err := os.Stat(filepath); err == nil {
		fmt.Println("File for " + concept + " already exists")
		return
	}

	resp, err := http.Get(url)
	handleError(err)
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	handleError(err)
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
