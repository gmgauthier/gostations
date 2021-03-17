package main

import (
	"fmt"
	"os"

	"github.com/dixonwille/wmenu/v5"
)

func Short( s string, i int ) string {
	runes := []rune( s )
	if len( runes ) > i {
		return string( runes[:i] )
	}
	return s
}


func Quit() {
	os.Exit(0)
}

func RadioMenu(stations []stationRecord) *wmenu.Menu {
	fmt.Println("...Radio Menu...")
	menu := wmenu.NewMenu("What is your choice?")
	menu.Action(
		func (opts []wmenu.Opt) error {
			if opts[0].Text == "Quit"{Quit()}
			val := fmt.Sprintf("%s",opts[0].Value)
			fmt.Printf("Streaming: " + opts[0].Text + "\n")
			stdout, _ := subExecute(player(), options(), val)
			fmt.Println(stdout)
			menu.Run()
			return nil
		})
	for _, station := range stations {
		stationListing := fmt.Sprintf("%-40s   %-5s %-5s %s", Short(station.Name, 40), station.Codec, station.Bitrate, station.Url)
		menu.Option(stationListing, station.Url, false, nil )
	}
	menu.Option("Quit", nil, true, nil)
	return menu
}