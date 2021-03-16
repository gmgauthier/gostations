package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alyu/configparser"
)

func str2int(strnum string) int {
	i, err := strconv.Atoi(strnum)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func Config(option string) (string, error) {
	configparser.Delimiter = "="
	configFile := "radiostations.ini"
	runtimeSection := "DEFAULT"

	config, err := configparser.Read(configFile)
	if err != nil {
		log.Fatal(err)
	}

	section, err := config.Section(runtimeSection)
	if err != nil {
		log.Fatal(err)
	}

	optval := section.Options()[option]
	if optval == "" {
		return "", errors.New("no value for option '%s'")
	}
	return optval, nil
}

func api() string {
	url, _ := Config("radio_browser.api")
	return url
}

func player() string {
	player, _ := Config("player.command")
	return player
}

func options() string {
	options, _ := Config("player.options")
	return options
}

func maxitems() int {
	items, _ := Config("menu_items.max")
	return str2int(items)
}