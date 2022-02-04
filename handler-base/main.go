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

func (b *Base) Setup() {
	b.Log = log.New()
	b.Params = &parameters.Params{}

	if len(b.Params.String("ProjectID")) < 1 {
		b.Log.Error("Handler Base : Setup was called without a ProjectID environment variable")
	}

	if len(b.Params.String("TimeZone")) < 1 {
		b.Log.Error("Handler Base : Setup was called without a TimeZone environment variable")
	}

	b.Db = fireStore.New(b.Params.String("ProjectID"))
	b.TimeZone = newTimeZone(b.Params.String("TimeZone"))
}