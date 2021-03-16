package main

import (
	"fmt"
	"github.com/dixonwille/wmenu/v5"
	"os"
)

func Quit() {
	os.Exit(0)
}

func RadioMenu(stations []stationRecord) *wmenu.Menu {
	menu := wmenu.NewMenu("What is your choice?")
	menu.Action(
		func (opts []wmenu.Opt) error {
			if opts[0].Text == "Quit"{Quit()}
			val := fmt.Sprintf("%s",opts[0].Value)
			fmt.Printf("Streaming: '" + opts[0].Text + "' at url - "+ val +"\n")
			stdout, _ := subExecute(player(), options(), val)
			fmt.Println(stdout)
			return nil
		})
	for _, station := range stations {
		menu.Option(station.Name, station.Url, false, nil )
	}
	menu.Option("Quit", nil, true, nil)
	return menu
}