package test1

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				time.Sleep(3 * time.Second)
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(30 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
