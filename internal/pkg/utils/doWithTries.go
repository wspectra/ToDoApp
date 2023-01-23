package utils

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	i := 1
	for j := attempts; j > 0; {
		log.Info().Msg(fmt.Sprintf("Trying to connect to database attempt %d (%d)", i, attempts))
		i += 1
		if err = fn(); err != nil {
			time.Sleep(delay)
			j--
			continue
		}
		return nil
	}
	return
}
