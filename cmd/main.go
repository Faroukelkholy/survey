package main

import (
	"faroukelkholy/survey/config"
	"log"
)

func main(){
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)
}
