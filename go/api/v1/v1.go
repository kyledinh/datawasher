package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kyledinh/datawasher/go/datastore"
	"github.com/kyledinh/datawasher/go/model"
	"github.com/kyledinh/datawasher/go/sys"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func NotAvailableJson(c *gin.Context) {
	c.JSON(401, gin.H{"message": "Required user authentication.", "status": sys.FAIL })
}

func UsageJson(c *gin.Context) {
	c.JSON(200, gin.H{"routes":"/washjson/"})
}

func GetRandomContact(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(datastore.Contacts))
	contact := datastore.Contacts[i]
	c.JSON(200, contact)
}

func GetContacts(c *gin.Context) {
	arr := datastore.Contacts
	log.Printf("... size of array object:  %v  ", len(arr))
	c.JSON(200, arr)
}

func PostWashJsonContacts(c *gin.Context) {
	rawbody, err := ioutil.ReadAll(c.Request.Body)
	//rawbody := []byte(`[{"first_name":"Ann","last_name":"Perkins","email":"ann.perkins@primeconsulting.com","phone_number":"555-207-1944","street_address":"72 Street Rd","city":"Pawnee","state":"IN"},{"first_name":"Donna","last_name":"Meagle","email":"donna.meagle@boeing.com","phone_number":"555-595-7884","street_address":"21 Street Rd","city":"Pawnee","state":"IN"}]`)
	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_READ_BODY, "status": sys.FAIL})
		return
	}
	//log.Printf("raw: % s \n", string(rawbody[:]))
	var arr []model.Contact
	json.Unmarshal(rawbody, &arr)

	for index := range arr {
		arr[index].First_name = datastore.RandFirstName()
		arr[index].Last_name = datastore.RandLastName()
	}
	c.JSON(200, arr)
}
