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

const RAND_FIRST_NAME = "MOX_RFN"
const RAND_LAST_NAME = "MOX_RLN"
const MOX_EMAIL = "MOX_EMAIL"
const RAND_STREET_ADDR = "MOX_RSA"
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

func ProcessAction(action string, str string) string {
    switch action {
    case RAND_FIRST_NAME:
        str = datastore.RandFirstName()
    case RAND_LAST_NAME:
        str = datastore.RandLastName()
    case RAND_STREET_ADDR:
        str = datastore.RandStreetAddress()
    }
    return str
}

func Setup() {
    SupportedActions = map[string]string {
        RAND_FIRST_NAME : "",
        RAND_LAST_NAME : "",
        MOX_EMAIL : "",
        RAND_STREET_ADDR : "",
    }
}
