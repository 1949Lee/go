package saver

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Got item #%d %v", itemCount, item)
		}
	}()

	return out
}
