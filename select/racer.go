package racer

import (
	"net/http"
	"time"
)

func Racer(urla, urlb string) (winner string) {
	aDuration := lapseTime(urla)
	bDuration := lapseTime(urlb)

	if aDuration < bDuration {
		return urla
	}

	return urlb
}

func lapseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
