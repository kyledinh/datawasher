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
