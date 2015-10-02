package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

var MACs = map[string]string{
	"Gatorade": "74:c2:46:3a:32:18",
}

func main() {
	client := &http.Client{}
	tcpdump := exec.Command("tcpdump", "-lveni", "any", "arp")
	stdout, err := tcpdump.StdoutPipe()
	if err != nil {
		fmt.Errorf("err: %s", err)
		return
	}

	if err := tcpdump.Start(); err != nil {
		fmt.Errorf("err: %s", err)
		return
	}

	fmt.Println("listening for arp from dash button")

	s := bufio.NewScanner(stdout)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		l := s.Text()
		isProbe := strings.Contains(l, "0.0.0.0")
		if isProbe {
		MAC:
			for name, mac := range MACs {
				if strings.Contains(l, mac) {
					fmt.Printf("%s Dash Pressed\n", name)
					err := switchLights(client)
					if err != nil {
						fmt.Errorf("err: %s", err)
					}
					break MAC
				}

			}
		}
	}
}

func switchLights(client *http.Client) error {
	lights := []string{"p3", "p4", "p17"}
	for _, light := range lights {
		req, err := http.NewRequest("GET", "http://10.0.0.13:1206/"+light, nil)
		if err != nil {
			return err
		}
		fmt.Println("making request", req)

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		fmt.Println("response Status:", resp.Status)
	}
	return nil
}
