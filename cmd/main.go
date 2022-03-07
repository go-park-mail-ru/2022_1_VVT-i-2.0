package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"serv"
)

func main() {
	configPath := flag.String("config", "../config/serv.toml", "path to config file")
	flag.Parse()

	servConfig := serv.NewConfig()
	_, err := toml.DecodeFile(*configPath, servConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := serv.Start(servConfig); err != nil {
		log.Fatal(err)
	}
}
