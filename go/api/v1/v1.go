package v1

import (
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/kyledinh/datawasher/go/datastore"
	"github.com/kyledinh/datawasher/go/model"
	"github.com/kyledinh/datawasher/go/sys"
	"github.com/kyledinh/datawasher/go/task"
	"github.com/kyledinh/datawasher/go/util"
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
		arr[index].Email = datastore.MakeEmailAddress(arr[index].First_name, arr[index].Last_name)
	}
	c.JSON(200, arr)
}

func PostWasher(c *gin.Context) {
	// task from query string
	log.Printf("c.Request.URL: %v", c.Request.URL)
	tasks := task.GetTasksFromURL(c.Request.URL)
	if tasks == nil {}

	// json payload
	rawbody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_READ_BODY, "status": sys.FAIL})
		return
	}
	rawbody = util.WrapJsonAsRoot(rawbody)
	log.Printf("IsRawbodyArray: % s \n", string(rawbody[:]))

	sj, err := simplejson.NewJson(rawbody)
	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_UNMARSHAL_BODY, "status": sys.FAIL})
		return
	}

	// process json payload with tasks packed in query string
	root := sj.Get("root")
	for i, _ := range root.MustArray() {
		for _, t := range tasks {
			root.GetIndex(i).Set(t.Field, task.ProcessAction(t.Action, ""))
		}
	}

	// 	TODO: root.GetIndex(i).Set("email", datastore.MakeEmailAddress(fn, ln))

	arr := root.MustArray()
	//log.Printf("... size of root array:  %v  ", len(arr))

	c.JSON(200, arr)
}
