package main

import (
	"log"
	"os"
	"path/filepath"
)

func createDir(fpath string) error {
	if err := os.Mkdir(filepath.Dir(fpath + string(filepath.Separator)), 0770); err != nil {
		return err
	}
	return nil
}

func createIniFile(fpath string) error {
	log.Printf("Creating config file: %s\n", fpath)
	if err := os.MkdirAll(filepath.Dir(fpath), 0770); err != nil {
		return err
	}

	file, err := os.Create(fpath)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	_, err = file.Write([]byte("[DEFAULT]\n"))
	_, err = file.Write([]byte("radio_browser.api=all.api.radio-browser.info\n"))
	_, err = file.Write([]byte("player.command=mpv\n"))
	_, err = file.Write([]byte("player.options=--no-video\n"))
	_, err = file.Write([]byte("menu_items.max=9999\n"))

	defer file.Close()

	return err
}
