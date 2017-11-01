package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kyledinh/datawasher/go/datastore"
	"github.com/kyledinh/datawasher/go/sys"
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
