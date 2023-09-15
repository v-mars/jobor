package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func SafeGoroutine(fn func()) {
	var err error
	go func() {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				log.Errorf("goroutine panic: %v", err)
			}
		}()
		fn()
	}()
}
