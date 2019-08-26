package main

import (
	"log"
	"github.com/aboglioli/coster-stocker/server/rest"
)

func main() {
	log.Println("[INTERFACES] Starting...")

	rest.Start();

	log.Println("[INTERFACES] Started")
}
