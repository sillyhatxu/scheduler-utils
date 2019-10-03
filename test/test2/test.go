package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Tick at for loop ", time.Now())
	}
	t := <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
	t = <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
	time.Sleep(7 * time.Second)
	t = <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
	t = <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
	t = <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
	t = <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
	t = <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
	t = <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
	t = <-ticker.C
	fmt.Println("Tick at", t, " --- ", time.Now())
}
