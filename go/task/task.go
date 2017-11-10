package task

import (
    "github.com/kyledinh/datawasher/go/datastore"
    "log"
    "net/url"
)

type Task struct {
    Field string
    Action string
}

const RAND_FIRST_NAME = "MOX_RFN"
const RAND_LAST_NAME = "MOX_RLN"
const RAND_FIRST_NAME_ESP = "MOX_RFN_ESP"
const RAND_LAST_NAME_ESP = "MOX_RLN_ESP"
const LIMIT = "limit"

func SupportedAction(field, action string) bool {
    if field == LIMIT {
        return false
    }
    return true
}

func GetTasksFromURL(u *url.URL) []Task {

    var fields []Task

    m := u.Query()
    for k, v := range m {
        log.Printf(" k: %v v: %v  ", k, v)

        if SupportedAction(k, v[0]) {
            task := Task{k, v[0]}
            fields = append(fields, task)
        }
	}
    //log.Printf("... size of tasks[] array:  %v  ", len(fields))
    return fields
}

func ProcessAction(action string, str string) string {
    switch action {
    case RAND_FIRST_NAME:
        str = datastore.RandFirstName()
    case RAND_LAST_NAME:
        str = datastore.RandLastName()
    }
    return str
}
