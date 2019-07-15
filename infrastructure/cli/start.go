package main

import (
	"fmt"
	"go-struct01/useCase"
	"os"
)

func main () {
	var fn = os.Getenv("fullname")
	var bd = os.Getenv("birthday")
	if len(fn) == 0  {
		fmt.Println("fullname is required and not be empty")
		os.Exit(1)
	}
	if len(bd) == 0  {
		fmt.Println("birthday is required and not be empry")
		os.Exit(1)
	}
	var e,r = useCase.SayHello(fn, bd)
	if e  != nil {
		fmt.Println(e)
		os.Exit(1)
	} else {
		fmt.Println(r)
		os.Exit(0)
	}

}

