package handlerBase

import (
	"github.com/driscollcode/log"
	"time"
)

func newTimeZone(timeZone string) *fetcher {
	return &fetcher{log: log.New()}
}

type fetcher struct {
	log  Log
	zone string
}

func (f *fetcher) Convert(now time.Time) time.Time {
	loc, err := time.LoadLocation(f.zone)
	if err != nil {
		f.log.Error("TimeZone Fetcher : Could not load time zone", ":", f.zone, ":", err.Error())
		return now
	}
	return now.In(loc)
}