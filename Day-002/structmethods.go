package main

import "fmt"

func main() {
	r := rect{
		width:  10,
		height: 20,
	}
	fmt.Println("The area is: ", r.area())
}

func (r rect) area() int {
	return r.width * r.height
}

type rect struct {
	width  int
	height int
}

// USER AUTHENTICATION
type authenticationDetails struct {
	username string
	password string
}

func (authD authenticationDetails) getBasicAuth() string {
	return fmt.Sprint(
		"Authorization: Basic %s:%s",
		authD.username,
		authD.password,
	)
}
