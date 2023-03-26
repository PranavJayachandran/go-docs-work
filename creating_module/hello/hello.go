package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
	log.SetPrefix("greetigns: ")
	log.SetFlags(0)

	names:=[]string{"Mohan","Sohan","Rohna"}

	messages,err:=greetings.Hellos(names)

	if err!=nil{
		log.Fatal(err)
	}
    fmt.Println(messages)
}