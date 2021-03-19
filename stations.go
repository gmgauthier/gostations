package main

import (
	"flag"
	"fmt"
	"os"
)

var version string

func showVersion(){
	fmt.Println(version)
}

func precheck(){
	if !isInstalled(player()){
		fmt.Printf("%s is either not installed, or not on your $PATH. Cannot continue.\n", player())
		os.Exit(1)
	}
}

func main(){
	argCount := len(os.Args[1:])

	var (
		name 	string
		country string
		state 	string
		tags 	string
		notok   bool
		version	bool
	)
	flag.Usage = func() {
		fmt.Printf("Usage: \n")
		fmt.Printf(" gostations ")
		fmt.Printf(" [-n \"name\"]  [-c \"home country\"]  [-s \"home state\"]  [-t \"ordered,tag,list\"] [-x] [-v]\n")
		flag.PrintDefaults()
		fmt.Printf("  -h (or none)\n")
		fmt.Printf("\tThis help message\n")
	}
	flag.StringVar(&name, "n", "", "Station name (or identifier).")
	flag.StringVar(&country, "c", "", "Home country.")
	flag.StringVar(&state, "s", "", "Home state (if in the United States).")
	flag.StringVar(&tags, "t", "", "Tag (or comma-separated tag list)")
	flag.BoolVar(&notok, "x", false,"If toggled, will show stations that are down")
	flag.BoolVar(&version, "v", false, "Show version.")
	flag.Parse()

	if argCount == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if version {
		showVersion()
		os.Exit(0)
	}

	precheck()

	stations, _ := StationSearch(name, country, state, tags, notok)
	menu := RadioMenu(stations)
	err := menu.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}