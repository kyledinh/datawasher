package datastore

import (
	"github.com/kyledinh/datawasher/go/cfg"
	"github.com/kyledinh/datawasher/go/model"
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func importContactsFromCSV() []model.Contact {
	csvFile, _ := os.Open(cfg.CSV_FILE)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var contacts []model.Contact
	var count int64 = 0
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		log.Printf(line[2])

		// SAMPLE CSV Entry
		// First Name,Last Name,Email,Phone Number,Street Address,City,State,Lifecycle Stage,Main Contact
		// Leslie,Knope,leslie.knope@cocaola.com,555-843-8116,68 Street Rd,Pawnee,IN,Lead,Phone
		if (count > 0) { // Ignore first line in CSV that has headers
			var c model.Contact

			c.First_name	= line[0]
			c.Last_name = line[1]
			c.Email = line[2]
			c.Phone_number = line[3]
			c.Street_address = line[4]
			c.City	= line[5]
			c.State = line[6]

			contacts = append(contacts, c)
		}
		count += 1
	}

	log.Printf("... slurped CSV File with %v entries ...", count)
	log.Printf("... number of records in Contacts %v", len(contacts))
	return contacts
}

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

func importNamesFromTxt(filename string) []string {
    var names []string
    file, err := os.Open(filename)
	if err != nil { log.Fatal(err) }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, strings.Title(scanner.Text()))
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
