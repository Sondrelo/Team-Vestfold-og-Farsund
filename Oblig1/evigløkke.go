package main

// Loopen teller uendelig oppover. For å stoppe loopen trykk ctrl + C, ctrl + break eller ctrl + delete.

import (
	"fmt"
	"os"
	"os/signal"

	//"time" // eller "runtime"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("STOP THIS!")
		os.Exit(1)
	}()

	//Loop som teller uendelig oppover til vi velger å avbrute/stoppe koden
	for i :=0;;i++ {
		fmt.Println(i)
	}


}
