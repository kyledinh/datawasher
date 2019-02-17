package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kyledinh/datawasher/go/api/v1"
	"github.com/kyledinh/datawasher/go/cfg"
	"github.com/kyledinh/datawasher/go/datastore"
	"github.com/kyledinh/datawasher/go/task"
	"log"
)

func main() {
	// SETUP
	cfg.Setup("cfg.json")
	task.Setup()
	datastore.Setup()
	portOption := ":" + cfg.HTTP_PORT

	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// AllowAllOrigins:  true,
	// AllowMethods:     []string{"GET", "POST"},
	// AllowHeaders:     []string{"Origin"},
	// ExposeHeaders:    []string{"Content-Length"},
	// AllowCredentials: true,
	// MaxAge: 12 * time.Hour,
	// }))

	// CORS MIDDLEWARE
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	// ROUTES
	router.GET("/", v1.UsageJson)
	router.GET("/random_contact", v1.GetRandomContact)
	router.GET("/contacts", v1.GetContacts)
	router.GET("/create", v1.GetCreate)
  router.GET("/reload", v1.ReloadDatastore)

	router.POST("/json_contacts", v1.PostWashJsonContacts)
	router.POST("/washer", v1.PostWasher)

	// SERVICES
	log.Printf("... started ServeFile ...")
	log.Printf("... using CORS allow all ...")
	log.Printf("... serving %v on port %v, LIMIT set to %v", cfg.APPNAME, cfg.HTTP_PORT, cfg.LIMIT)
	if cfg.HTTP_PORT == "443" {
		router.RunTLS(portOption, "certs/cert.pem", "certs/key.pem")
	} else {
		router.Run(portOption)
	}
}
