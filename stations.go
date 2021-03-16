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
	)

	flag.StringVar(&name, "n", "", "Station name (or identifier).")
	flag.StringVar(&country, "c", "", "Home country.")
	flag.StringVar(&state, "s", "", "Home state (if in the United States).")
	flag.StringVar(&tags, "t", "", "Tag (or comma-separated tag list)")
	flag.Parse()

	if argCount == 0 {
		flag.Usage()
	}

	stations, err := GetStations("tag=chicago")
	if err != nil{
		fmt.Println(err.Error())
	}
	for _, station := range stations {
		fmt.Printf("\"%s\", %s, %s, %s\n", station.Name, station.Codec, station.Bitrate, station.Url)
	}


}