package helloC

import "fmt"

import (
	proto "example.com/mP"
)

func Id() string { return "Client" }
func Hello() { fmt.Printf("%s <= %s\n", Id(), proto.Id()) }

