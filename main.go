package main

import (
	"log"
	"os"
	"strings"

	"github.com/ProtonMail/go-autostart"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const APP_VERSION = "1.0"

func main() {
	app := kingpin.New("proxy", "happy with proxy")
	app.Author("snail").Version(APP_VERSION)
	nameKey := app.Flag("key", "Unique identifier for the application").Short('k').Default("").String()
	name := app.Flag("name", "The app name").Short('n').Default("").String()
	enable := app.Command("enable", "enable application auto startup")
	command := enable.Flag("command", "command of auto startup").Short('c').Default("").String()
	app.Command("disable", "disable application auto startup")
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	if strings.Trim(*nameKey, " ") == "" {
		app.Fatalf("namd key is required")
		return
	}
	if cmd == "enable" {
		if strings.Trim(*command, " ") == "" {
			app.Fatalf("command is required")
			return
		}
	}
	if *name == "" {
		*name = *nameKey
	}
	log.Println(*command, strings.Fields(*command))
	autostart := &autostart.App{
		Name:        *nameKey,
		DisplayName: *name,
		Exec:        strings.Fields(*command),
	}

	if cmd == "enable" {
		if autostart.IsEnabled() {
			log.Println(*name + " is already enabled")
			return
		}
		log.Println("Enabling app...")
		if err := autostart.Enable(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := autostart.Disable(); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Done!")
}
