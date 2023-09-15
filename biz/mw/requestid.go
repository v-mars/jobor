package mw

import (
	"github.com/sirupsen/logrus"
)

type RequestIdHook struct{}

func (h *RequestIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *RequestIdHook) Fire(e *logrus.Entry) error {
	ctx := e.Context
	if ctx == nil {
		return nil
	}
	value := ctx.Value("X-Request-Id")
	if value != nil {
		e.Data["log_id"] = value
	}
	return nil
}

func InitRequestid() {
	//logger := hertzlogrus.NewLogger(hertzlogrus.WithHook(&RequestIdHook{}))
	//hlog.SetLogger(logger)

}
