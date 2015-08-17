package main

import (
       "net/http"
       "log"
       "github.com/stianeikeland/go-rpio"
       "io"
       "fmt"
)

var p2, p3, p4, p17 bool

var html = `
<html>
<body>
<h3>Turkish Lamp: %s</h3>
<a href="/p2">Toggle Turkish Lamp</a>
<h3>Desk Lamp: %s</h3>
<a href="/p3">Toggle Desk Lamp</a>
<h3>Tall Lamp: %s</h3>
<a href="/p4">Toggle Tall Lamp</a>
<h3>Fan: %s</h3>
<a href="/p17">Toggle Fan</a>
</body>
</html>
`
var (
    // Use mcu pin 2, corresponds to physical pin 3 on the pi
    pin2 = rpio.Pin(2)
    // Use mcu pin 3, corresponds to physical pin 5 on the pi
    pin3 = rpio.Pin(3)		  
    // Use mcu pin 4, corresponds to physical pin 7 on the pi
    pin4 = rpio.Pin(4)		  
    // Use mcu pin 17, corresponds to physical pin 11 on the pi
    pin17 = rpio.Pin(17)		  

)

func stateString(state bool) string {
     if state {
     	return "on"
     } 
     return "off"
}

func Pin2Switch(w http.ResponseWriter, req *http.Request) {
     pin2.Toggle()
     p2 = !p2
     io.WriteString(w, fmt.Sprintf(html, stateString(p2), stateString(p3), stateString(p4), stateString(p17)))
}

func Pin3Switch(w http.ResponseWriter, req *http.Request) {
     pin3.Toggle()
     p3 = !p3
     io.WriteString(w, fmt.Sprintf(html, stateString(p2), stateString(p3), stateString(p4), stateString(p17)))
}

func Pin4Switch(w http.ResponseWriter, req *http.Request) {
     pin4.Toggle()
     p4 = !p4
     io.WriteString(w, fmt.Sprintf(html, stateString(p2), stateString(p3), stateString(p4), stateString(p17)))
}

func Pin17Switch(w http.ResponseWriter, req *http.Request) {
     pin17.Toggle()
     p17 = !p17
     io.WriteString(w, fmt.Sprintf(html, stateString(p2), stateString(p3), stateString(p4), stateString(p17)))
}

func Main(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, fmt.Sprintf(html, stateString(p2), stateString(p3), stateString(p4), stateString(p17)))
}

func main() {
// Open and map memory to access gpio, check for errors
   if err := rpio.Open(); err != nil {
      fmt.Println(err)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pins to output mode
	pin2.Output()
	pin3.Output()	
	pin4.Output()	
	pin17.Output()	
     http.HandleFunc("/p2", Pin2Switch)
     http.HandleFunc("/p3", Pin3Switch)
     http.HandleFunc("/p4", Pin4Switch)
     http.HandleFunc("/p17", Pin17Switch)
     http.HandleFunc("/main", Main)
     http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
     fmt.Printf("Serving %s on HTTP port: %d\n", ".", 1206)
     log.Fatal(http.ListenAndServe(":1206", nil))
}
