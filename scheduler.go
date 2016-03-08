package main

import "time"

func Scheduler() {
	go func() {
		c := time.Tick(6 * time.Minute)
		for range c {
			// Note this purposfully runs the function
			// in the same goroutine so we make sure there is
			// only ever one. If it might take a long time and
			// it's safe to have several running just add "go" here.
			UpdateHashtags()
		}
	}()

	// Or to block forever:
	select {}
}
