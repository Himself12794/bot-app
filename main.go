package main

import (
	"log"

	server "github.com/Himself12794/bot-app/batcave"
)

const (
	apiURL    = "https://api.ciscospark.com/v1/messages"
	testRoom  = "Y2lzY29zcGFyazovL3VzL1JPT00vNWZmNGM1ZTAtZjNhMS0xMWU2LWJhOWYtOTUwN2UyMTZkOTRj"
	testToken = "MzVhMzc3NzYtNDNjYS00MWZkLWJjODgtN2JjMWIwNzgzYTY4YjMwZjE4MGMtNGFj"
)

func main() {

	log.Println("Starting server")
	server.Start()
}
