package datastore

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"github.com/kyledinh/datawasher/go/model"
	"github.com/kyledinh/datawasher/go/cfg"
)

var Contacts []model.Contact

func GetTest (num int) ([]model.Contact) {
	if (num > len(Contacts)) {
		num = len(Contacts)
	}
	var arr []model.Contact
	for i := 0; i < num; i += 1 {
		arr = append(arr, Contacts[i])
	}
	return arr
}

func Setup() {
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
	Contacts = contacts
	log.Printf("... slurped CSV File with %v entries ...", count)
	log.Printf("... number of records in Contacts %v", len(Contacts))
}
