package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Configuration values:")
	fmt.Println("api url: ", api())
	fmt.Println("player command: ", player())
	fmt.Println("player options: ", options())
	fmt.Println("maximum menu items: ", maxitems())

	argCount := len(os.Args[1:])

	var (
		name 	string
		country string
		state 	string
		tags 	string
	)

	flag.StringVar(&name, "n", "", "Station name (or identifier).")
	flag.StringVar(&country, "c", "", "Home country.")
	flag.StringVar(&state, "s", "", "Home state (if in the United States).")
	flag.StringVar(&tags, "t", "", "Tag (or comma-separated tag list)")
	flag.Parse()

	if argCount == 0 {
		flag.Usage()
	}

}