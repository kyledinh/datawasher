package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// defaults
var APPNAME string = "appname"
var HTTP_PORT string = "8000"
var SERVER string = "0.0.0.0"
var CSV_FILE string = "data/contacts.csv"
var LIMIT int = 10000
var VERSION string = "0.0"

var conf JsonCfg

type JsonCfg struct {
	Appname   string
	Http_port string
	Server    string
	Csv_file  string
	Limit     int
  Appversion   string
}

// ConfigFrom is a method to read json properties
func (cfg *JsonCfg) ConfigFrom(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &cfg)
	if err != nil {
		log.Printf("bad json", err)
	}

	return
}

// ConfigSetup is the init function called from app.go, to read the cfg.json file
// If error reading the cfg.json; defaults are used
func Setup(path string) {
	err := conf.ConfigFrom(path)
	if err != nil {
		log.Printf("######### CRITICAL: Failed to load cfg.json %v", err)
		log.Printf("######### USING DEFAULT CFG SETTINGS!!!")
	} else {
		APPNAME = conf.Appname
		HTTP_PORT = conf.Http_port
		SERVER = conf.Server
		CSV_FILE = conf.Csv_file
		LIMIT = conf.Limit
    VERSION = conf.Appversion
		log.Printf("... loaded cfg.json ...")
		log.Printf("CSV_FILE: %v", CSV_FILE)
    log.Printf("Version: %v", VERSION)
	}
}
