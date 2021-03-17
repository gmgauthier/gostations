package main

import (
	"log"
	"os"
	"path/filepath"
)

func closeFile(file *os.File, errorlist []error) {
	err := file.Close()
	if err != nil {
		errorlist = append(errorlist, err)
	}
}

func createIniFile(fpath string) []error {
	log.Printf("Creating config file: %s\n", fpath)
	var errorlist []error
	if err := os.MkdirAll(filepath.Dir(fpath), 0770); err != nil {
		errorlist = append(errorlist, err)
	}

	file, err := os.Create(fpath)
	if err != nil {
		errorlist = append(errorlist, err)
	}

	if _, err := file.Write([]byte("[DEFAULT]\n")); err != nil {
		errorlist = append(errorlist, err)
	}
	if _, err := file.Write([]byte("radio_browser.api=all.api.radio-browser.info\n")); err != nil {
		errorlist = append(errorlist, err)
	}
	if _, err := file.Write([]byte("player.command=mpv\n")); err != nil {
		errorlist = append(errorlist, err)
	}
	if _, err := file.Write([]byte("player.options=--no-video\n")); err != nil {
		errorlist = append(errorlist, err)
	}
	if _, err := file.Write([]byte("menu_items.max=9999\n")); err != nil {
		errorlist = append(errorlist, err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			errorlist = append(errorlist, err)
		}
	}()

	return errorlist
}
