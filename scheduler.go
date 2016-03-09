package main

import "time"

func Scheduler() {
	c := time.Tick(6 * time.Minute)
	for range c {
		UpdateHashtags()
	}

	// To block forever:
	select {}
}
