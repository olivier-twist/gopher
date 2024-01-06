package main

import (
	"fmt"
)

func main() {
	u := User.CreateNew("Jack", "Damon")
	fmt.Println(u.String())
}
