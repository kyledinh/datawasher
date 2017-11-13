package task

import (
    "github.com/kyledinh/datawasher/go/datastore"
    "log"
    "net/url"
    "strconv"
)

type Task struct {
    Field string
    Action string
}

type Setting struct {
    First_name string
    Last_name string
    Email string
    Limit int
}

// Actions
const MOX_EMAIL = "MOX_EMAIL"
const RAND_FIRST_NAME = "MOX_RFN"
const RAND_LAST_NAME = "MOX_RLN"
const RAND_STREET_ADDR = "MOX_RSA"
const RAND_INT_100 = "MOX_RI_100"
const RAND_INT_1000 = "MOX_RI_1000"
const RAND_PHONE_555 = "MOX_P555"
// Field
const LIMIT = "limit"

var SupportedActions map[string]string

func SupportedAction(field, action string) bool {
    if _, ok := SupportedActions[action]; ok {
        return true
    }
    return false
}

func GetTasksAndSettings(u *url.URL) ([]Task, Setting) {
    var tasks []Task
    var setting Setting
    setting.Limit = 10

    m := u.Query()
    for k, v := range m {
        log.Printf(" k: %v v: %v  ", k, v)

        if SupportedAction(k, v[0]) {
            task := Task{k, v[0]}
            tasks = append(tasks, task)
        }
        if k == LIMIT {
            limit, err := strconv.Atoi(v[0])
            if err == nil {
                setting.Limit = limit
            }
        }
	}

    for _, t := range tasks {
        if t.Action == RAND_FIRST_NAME { setting.First_name = t.Field }
        if t.Action == RAND_LAST_NAME { setting.Last_name = t.Field }
        if t.Action == MOX_EMAIL { setting.Email = t.Field }
    }
    return tasks, setting
}

func ProcessAction(action string) interface{} {
    var val interface{}
    switch action {
    case RAND_FIRST_NAME:
        val = datastore.RandFirstName()
    case RAND_LAST_NAME:
        val = datastore.RandLastName()
    case RAND_STREET_ADDR:
        val = datastore.RandStreetAddress()
    case RAND_INT_100:
        val = datastore.RandInt(100)
    case RAND_INT_1000:
        val = datastore.RandInt(1000)
    }
    return val
}

func Setup() {
    SupportedActions = map[string]string {
        MOX_EMAIL : "",
        RAND_FIRST_NAME : "",
        RAND_LAST_NAME : "",
        RAND_STREET_ADDR : "",
        RAND_INT_100 : "",
        RAND_INT_1000 : "",
        RAND_PHONE_555 : "",
    }
}
