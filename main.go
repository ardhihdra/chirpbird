package main

import (
	"os"

	"github.com/ardhihdra/chirpbird/chat"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	if port == "" {
		port = "4000"
	}
	chat.Start(port)
}
