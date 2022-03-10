package main

import (
	"flag"
	"log"

	// "serv"

	"github.com/BurntSushi/toml"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv"
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
