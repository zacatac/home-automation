package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

var (
	// Use mcu pin 2, corresponds to physical pin 3 on the pi
	pin = rpio.Pin(2)
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()
	// http.HandleFunc("/hello", HelloServer)
	// http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	// fmt.Printf("Serving %s on HTTP port: %d\n", ".", 1206)
	// log.Fatal(http.ListenAndServe(":1206", nil))
	pin.Toggle()
	time.Sleep(2 * time.Second)
	pin.Toggle()	
}
