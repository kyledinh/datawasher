package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kyledinh/datawasher/go/datastore"
	"github.com/kyledinh/datawasher/go/sys"
	"log"
)

func NotAvailableJson(c *gin.Context) {
	c.JSON(401, gin.H{"message": "Required user authentication.", "status": sys.FAIL })
}

func UsageJson(c *gin.Context) {
	c.JSON(200, gin.H{"routes":"/washjson/"})
}

func GetWashJson(c *gin.Context) {
	c.JSON(200, gin.H{"name_first":"Bob", "last_name":"Lackey"})
}

func GetContacts(c *gin.Context) {
	arr := datastore.Contacts
	log.Printf("... size of array object:  %v  ", len(arr))
	c.JSON(200, arr)

}
