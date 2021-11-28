package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ardhihdra/chirpbird/chat"
	"github.com/ardhihdra/chirpbird/controllers"
)

var (
	port  = os.Getenv("PORT")
	users = controllers.NewUsersController()
)

func main() {
	if port == "" {
		port = "4000"
	}
	log.Printf("listening to %s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from chirpbird!"))
	})

	http.HandleFunc("/login", users.Login())

	http.HandleFunc("/Register", users.Register())

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

	chat.Start(port)
}
