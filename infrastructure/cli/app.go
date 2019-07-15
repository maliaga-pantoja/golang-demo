package main

import(
	"fmt"
	"go-struct01/useCase"
	"os"
)

func  main()  {
	var fullname = os.Getenv("fullname")
	var birthday = os.Getenv("birthday")
	var useCaseOk,useCaseError = useCase.SayHello(fullname, birthday)
	if  useCaseError != nil {
		fmt.Print(useCaseError)
	} else {
		fmt.Println(useCaseOk)
	}
}