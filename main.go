package main

import (
	"fmt"

	"github.com/olivier-twist/gopher/user"
)

func main() {
	u := user.User{}
	u.CreateNew("Jack", "Famon")
	fmt.Println(u.String())
}
