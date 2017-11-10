package datastore

import (
	"bufio"
	"log"
	"os"
)

func importEmailDomains() []string {
    var arr []string
    file, err := os.Open("data/domains.txt")
	if err != nil { log.Fatal(err) }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	if err := scanner.Err(); err != nil { log.Fatal(err) }
    return arr
}

func importFirstNames() []string {
    var names []string
    file, err := os.Open("data/first-names.txt")
	if err != nil { log.Fatal(err) }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	if err := scanner.Err(); err != nil { log.Fatal(err) }
    return names
}

func importLastNames() []string {
    var names []string
    file, err := os.Open("data/last-names.txt")
	if err != nil { log.Fatal(err) }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	if err := scanner.Err(); err != nil { log.Fatal(err) }
    return names
}

func importStates() []string {
	states := []string{ "AL","AK","AZ","AR","CA","CO","CT","DE",
		"FL","GA","HI","ID","IL","IN","IA","KS","KY","LA","ME","MD",
		"MA","MI","MN","MS","MO","MT","NE","NV","NH","NM","NY","NC",
		"ND","OH","OK","OR","PA","RI","SC","SD","TN","TX","UT","VT",
		"VA","WA","WV","WI","WY",
	}
	return states
}

func importStreetNames() []string {
	streets := []string{
		"Elm","Maple","Oak","Spruce","Birch","Forest","Park",
		"Washington","Lincoln","Adams","Jefferson","Hayes",
		"Main","Broad","Central","Market","Post",
	}
	return streets
}

func importStreetTypes() []string {
	return []string {
		"Ave", "Avenue", "St", "Street", "Blvd", "Alley", "Way", "Road",
		"Dr", "Ct", "Court", "Pl", "Terrace", "Loop",
	}
}
