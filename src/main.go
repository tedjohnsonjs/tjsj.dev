package main

import (
	"log"
	"github.com/tedski999/tjsj.dev/src/content"
	"github.com/tedski999/tjsj.dev/src/server"
)

func main() {
	log.Println("Loading content...")
	content.LoadTemplates("./templates/*.html")
	content.LoadTitleGoofs("./data/titlegoofs.txt")
	log.Println("Starting HTTPS server...")
	log.Fatal(server.StartServer())
}
