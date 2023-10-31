package utils

import (
	"github.com/robfig/cron/v3"
	"time"
)

type WrapCron struct {
	cr *cron.Cron
	C  chan time.Time
}

func NewWrapCron(spec string) (*WrapCron, error) {
	wc := &WrapCron{
		cr: cron.New(cron.WithSeconds()),
		C:  make(chan time.Time),
	}
	if _, err := wc.cr.AddFunc(spec, func() {
		wc.C <- time.Now()
	}); err != nil {
		return nil, err
	}

	return wc, nil

}
func (wc *WrapCron) Start() {
	wc.cr.Start()
}
