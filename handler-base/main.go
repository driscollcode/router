package handlerBase

import (
	fireStore "github.com/driscollcode/firestore"
	"github.com/driscollcode/log"
	"github.com/driscollcode/parameters"
)

type Base struct {
	Log      Log
	Db       Database
	Params   Params
	TimeZone Timezone
}

func (b *Base) Setup(projectID, timeZone string) {
	b.Log = &log.Log{}
	b.Params = &parameters.Params{}
	b.Db = fireStore.New(projectID)
	b.TimeZone = newTimeZone(timeZone)
}