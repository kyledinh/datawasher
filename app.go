package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kyledinh/datawasher/go/api/v1"
	"github.com/kyledinh/datawasher/go/cfg"
	"github.com/kyledinh/datawasher/go/datastore"
	"log"
)

func main() {
	// SETUP
	cfg.Setup("cfg.json")
	datastore.Setup()
	portOption := ":" + cfg.HTTP_PORT

	// ROUTES
	g := gin.Default()
	g.GET("/", v1.UsageJson)
	g.GET("/random_contact", v1.GetRandomContact)
	g.GET("/contacts", v1.GetContacts)

	g.POST("/json_contacts", v1.PostWashJsonContacts)

	// SERVICES
	log.Printf("... started ServeFile ...")
	log.Printf("... serving %v on port %v ", cfg.APPNAME, cfg.HTTP_PORT)
	g.Run(portOption);
}
