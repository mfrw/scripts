package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

/* A better solution here is that Hard Code the
username and pass, then build, it would be difficult
to get the password that way
*/

func main() {
	resp, err := http.Get("http://www.google.com/")
	if err != nil {
		fmt.Println("Error encountered")
	}
	var username string
	var password string
	flag.StringVar(&username, "u", "falak16018", "Your username")
	flag.Parse()

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	finalURL := resp.Request.URL.String()
	bdy := string(body)
	if strings.Contains(bdy, "IIIT-D") {
		fmt.Print("Password: ")
		terminalEcho(false)
		fmt.Scanln(&password)
		terminalEcho(true)
		fmt.Println("")
		http.PostForm(finalURL, url.Values{"username": {username}, "magic": {bdy[6354:6370]}, "password": {password}})
		//		fmt.Println("Logged in successfully")
	} else {
		fmt.Println("Already Logged in dude")
	}
}

func terminalEcho(show bool) {
	// Enable or disable echoing terminal input. This is useful specifically for
	// when users enter passwords.
	// calling terminalEcho(true) turns on echoing (normal mode)
	// calling terminalEcho(false) hides terminal input.
	var termios = &syscall.Termios{}
	var fd = os.Stdout.Fd()

	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd,
		syscall.TCGETS, uintptr(unsafe.Pointer(termios))); err != 0 {
		return
	}

	if show {
		termios.Lflag |= syscall.ECHO
	} else {
		termios.Lflag &^= syscall.ECHO
	}

	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd,
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(termios))); err != 0 {
		return
	}
}
