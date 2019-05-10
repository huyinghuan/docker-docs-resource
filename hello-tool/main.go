package main

import (
	"flag"
	"fmt"
	"time"
)

var timeInfo = flag.Bool("t", false, "show date info")
var msg = flag.String("m", "", "show meessage")
var name = flag.String("n", "tool", "show name")

func main() {

	flag.Parse()
	fmt.Printf("I'm a %s, I get this args:\n", *name)
	if *timeInfo {
		fmt.Println(time.Now().Local())
	}
	if *msg != "" {
		fmt.Println(*msg)
	}
}
