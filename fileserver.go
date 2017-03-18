package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	/* A static web server serving a current path
	It can also be used a hacky file tx program.
	Just build it with with the arch you want to
	run on, a static ready to run binary is the output:

	GOOS=windows go build fileserver.go // for windows
	GOOS=openbsd go build fileserver.go // for openbsd
	GOOS=linux   go build fileserver.go // for linux
	*/

	path := ""
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [path]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Using current dir as path\n\n")
		tpath, err := os.Getwd() // get the current working dir
		if err != nil {
			fmt.Fprintf(os.Stderr, "Some bad stuff occured\n")
			os.Exit(-1)
		}
		path = tpath
	} else {
		path = os.Args[1]
	}

	/* get IP Addrs of Interfaces */
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get Interfaces\n")
		os.Exit(-1)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get Addrs\n")
			os.Exit(-1)
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Println(ip.String())
			// process IP address
		}
	}
	/* done getting ip addrs */

	fmt.Println("Try any of those ipaddrs with xx.xx.xx.xx:8080")
	fmt.Println("Server starting on Port 8080")
	fmt.Println("Server Hosting: ", path)
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(path))))
}
