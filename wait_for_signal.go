package utils

import (
	"os"
	"os/signal"
)

// Блокировка в ожидании сигнала.
func WaitForSignal(signals ...os.Signal) os.Signal {
	var c = make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	return <-c
}
