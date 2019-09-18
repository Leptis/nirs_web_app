package main

import (
	"log"
	"nirs_web_app/internal/app/server"
)

func main(){
	s := server.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
