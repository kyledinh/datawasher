package datastore

import (
	//"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"github.com/kyledinh/datawasher/go/model"
)

var Contacts []model.Contact
var FirstNames []string
var FemaleNames []string
var MaleNames []string
var LastNames []string
var EmailDomains []string
var States []string
var StreetNames []string
var StreetTypes []string

func RandFirstName (sex string) string {
	timeseed := time.Now().UnixNano()
	rand.Seed(timeseed)
	if (sex == "M") {
		i := rand.Intn(len(MaleNames))
		return MaleNames[i]
	}
	if (sex == "F") {
		i := rand.Intn(len(FemaleNames))
		return FemaleNames[i]
	}
	i := rand.Intn(len(FirstNames))
	return FirstNames[i]
}

func RandLastName () string {
	timeseed := time.Now().UnixNano()
	rand.Seed(timeseed)
	i := rand.Intn(len(LastNames))
	return LastNames[i]
}

func MakeEmailAddress (first, last string) string {
	user := "user"
	first = strings.ToLower(first)
	last = strings.ToLower(last)
	i := rand.Intn(len(EmailDomains))

	switch rand.Intn(4) {
	case 0:
		user = first + last
	case 1:
		user = first[:1] + last
	case 2:
		user = first + last[:1]
	default:
		user = first + "." + last
	}
	user = strings.Replace(user, " ", "", -1)
	user = strings.Replace(user, "'", "", -1)
	return user + "@" + EmailDomains[i]
}

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

func RandInt(limit int) int {
	return rand.Intn(limit)
}

func RandStateCode() string {
	i := rand.Intn(len(States))
	return States[i]
}

func streetNumber() string {
	i := rand.Intn(3000)
	return strconv.Itoa(i)
}

func RandStreetAddress() string {
	i := rand.Intn(len(StreetNames))
	j := rand.Intn(len(StreetTypes))
	return streetNumber() + " " + StreetNames[i] + " " + StreetTypes[j]
}

func RandPhoneNumber() string {
	// TODO
	return "415-555-0123"
}

func RandSexMF() string {
	i := rand.Intn(100)
	if i % 2 ==  0 {
		return "M"
	}
	return "F"
}

func Setup() {
	Contacts = importContactsFromCSV()
	FirstNames = importNamesFromTxt("data/first-names.txt")
	FemaleNames = importNamesFromTxt("data/female-names.txt")
	MaleNames = importNamesFromTxt("data/male-names.txt")
	LastNames = importNamesFromTxt("data/last-names.txt")
	EmailDomains = importEmailDomains()
	States = importStates()
	StreetNames = importStreetNames()
	StreetTypes = importStreetTypes()
}
