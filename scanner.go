package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var gotcha [23]string

func main() {
	wg := sync.WaitGroup{}

	content, err := ioutil.ReadFile("target.txt")
	if err != nil {
		log.Fatal(err)
	}
	target := string(content)

	const firstport int = 1
	const lastport int = 65535
	const threads int = 65535
	const porttimeout int = 5

	if target == "" {
		os.Exit(4)
	}

	fmt.Println("Starting scan. Open ports will be listed below.\n---------------------------------------------------\n\nPORT	SERVİCE\n----	-------")
	var a int = 0
	for i := firstport; i <= lastport; i++ {
		targetHostWithPort := fmt.Sprintf("%s:%d", target, i)
		port := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isTcpPortOpen(targetHostWithPort, porttimeout) {
				servicequery(port)
				fmt.Println("")
				gotcha[a] = strconv.Itoa(port)
				a = a + 1

			}

		}()

	}

	var openportstr string
	for say := 0; say <= a; say++ {
		openportstr += gotcha[say]
		if say < a-1 {
			openportstr += ","
		}

	}
	wg.Wait()
	nmapscantrue(string(target), openportstr)
	println("")

}

func isTcpPortOpen(targetHostWithPort string, portTimeout int) bool {
	_, err := net.DialTimeout("tcp", targetHostWithPort, time.Second*time.Duration(portTimeout))
	if err == nil {
		return true
	}
	return false
}

func servicequery(port int) {
	if port == 20 {
		fmt.Printf("%d	FTP-DATA", port)

	} else if port == 21 {
		fmt.Printf("%d	FTP", port)
	} else if port == 22 {
		fmt.Printf("%d	SSH", port)
	} else if port == 23 {
		fmt.Printf("%d	TELNET", port)
	} else if port == 25 {
		fmt.Printf("%d	SMTP", port)
	} else if port == 53 {
		fmt.Printf("%d	DNS", port)
	} else if port == 80 {
		fmt.Printf("%d	HTTP", port)
	} else if port == 443 {
		fmt.Printf("%d	HTTPS", port)
	} else if port == 161 || port == 162 {
		fmt.Printf("%d	SNMP", port)
	} else if port == 3389 {
		fmt.Printf("%d	REMOTE DESKTOP", port)
	} else {
		fmt.Printf("%d	UNKOWN", port)
	}

}

func nmapscantrue(target string, openportstr string) {

	fmt.Println("\n---------------------------------------------------\n Waiting for NMAP output \n--------------------------------------------------- ")
	var nmapcommand string = "nmap " + target + " -p " + openportstr + " -O -A -sA -sV -T5 -Pn"

	e := os.Remove("command.txt")
	if e != nil {
		file, err := os.OpenFile("command.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(nmapcommand)
	} else {

		file, err := os.OpenFile("command.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(nmapcommand)

	}



}

