package batcave

//import (
//    "encoding/gob"
//)

const file = "/tmp/tasks.gob"

// TaskDatabase is an alias for a map, where the key is userID, and value is list of tasks 
type TaskDatabase map[string][]string

// NewTaskDatabase creates an new object for this type
func NewTaskDatabase() TaskDatabase {
    return make(map[string][]string)
}

// AddTaskForUser adds a new task for the specified user.
// Will create the user if it does not exist.
func (t TaskDatabase) AddTaskForUser(user string, task string) {
    
    var tasks []string
    
    if val, ok := t[user]; ok {
        tasks = val
    } 
    
    t[user] = append(tasks, task)
}
