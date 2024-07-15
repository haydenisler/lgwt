package kselect

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimer = 10 * time.Second

func Racer(u1, u2 string) (winner string, err error) {
	return ConfigurableRacer(u1, u2, tenSecondTimer)
}

func ConfigurableRacer(u1, u2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(u1):
		return u1, nil
	case <-ping(u2):
		return u2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out wating for %s and %s", u1, u2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
