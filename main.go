package main

import (
	"github.com/Kangrao0o/goepub/config"
	"log"
)

func main() {
	conf := config.Get()
	log.Println("conf:", conf)
}
