package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	argCount := len(os.Args[1:])

	var (
		name 	string
		country string
		state 	string
		tags 	string
		notok   bool
	)
	flag.Usage = func() {
		fmt.Printf("Usage: \n")
		fmt.Printf(" gostations ")
		fmt.Printf(" [-n \"name\"]  [-c \"home country\"]  [-s \"home state\"]  [-t \"ordered,tag,list\"] [-x]\n")
		flag.PrintDefaults()
		fmt.Printf("  -h (or none)\n")
		fmt.Printf("\tThis help message\n")
	}
	flag.StringVar(&name, "n", "", "Station name (or identifier).")
	flag.StringVar(&country, "c", "", "Home country.")
	flag.StringVar(&state, "s", "", "Home state (if in the United States).")
	flag.StringVar(&tags, "t", "", "Tag (or comma-separated tag list)")
	flag.BoolVar(&notok, "x", false,"If toggled, will show stations that are down")
	flag.Parse()

	if argCount == 0 {
		flag.Usage()
		os.Exit(0)
	}

	stations, _ := StationSearch(name, country, state, tags, notok)

	menu := RadioMenu(stations)
	err := menu.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}