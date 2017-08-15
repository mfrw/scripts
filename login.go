package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// Gets the credentials from the user and returs them as strings
func getCreds() (string, string) {
	var username string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("")
	return username, string(password)
}

// Login takes in username and password and attempts to login
// If successfull returns the logout URL
func Login(username, password string) (string, error) {
	res, err := http.Get("http://www.google.com/")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.Request.URL.Hostname() != "auth.iiitd.edu.in" {
		return "", fmt.Errorf("Already Connected")
	}
	// get magic from URL and reply URL
	magic := res.Request.URL.RawQuery
	u := res.Request.URL.String()

	// pack data
	data := url.Values{
		"username": {username},
		"magic":    {magic},
		"password": {password},
	}

	resp, err := http.PostForm(u, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body[4816:4870]), nil

}

//  TODO: Add functionality to get password from a file //

func main() {
	username, password := getCreds()

	logout, err := Login(username, password)
	if err != nil {
		log.Println("Could not login", err)
	} else {
		fmt.Println("LogoutURL: ", logout)
	}
}
