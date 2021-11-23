package handlerBase

import (
	FireStore "cloud.google.com/go/firestore"
	fireStore "github.com/driscollcode/firestore"
	"time"
)

type Log interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Notice(msg ...interface{})
	Error(msg ...interface{})
	Alert(msg ...interface{})
}

type Database interface {
	Delete(filename string) error
	Fetch(filename string, destination interface{}) error
	FetchAll(path string, template interface{}) (map[string]interface{}, error)
	Search(path string, queries ...fireStore.Query) ([]*FireStore.DocumentSnapshot, error)
	SearchOne(container interface{}, path string, queries ...fireStore.Query) error
	SearchOneRaw(path string, queries ...fireStore.Query) (map[string]interface{}, error)
	Update(path string, values map[string]interface{}) error
	Write(path string, object interface{}) error
}

type Params interface {
	Bool(param string) bool
	Int(param string) int
	String(param string) string
	StringWithDefault(param, defaultVal string) string
}

type Timezone interface {
	Convert(now time.Time) time.Time
}
