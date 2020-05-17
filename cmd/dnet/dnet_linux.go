package main

import (
	"os"
	"os/signal"
	"syscall"

	psignal "github.com/demonoid81/moby/pkg/signal"
)

func setupDumpStackTrap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			psignal.DumpStacks("")
		}
	}()
}
