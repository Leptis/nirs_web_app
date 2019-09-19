package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"nirs_web_app/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "congif-path", "configs/server.toml", "path to server config file")
}

func main(){
	flag.Parse()
	config := server.NewConfig()

	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
