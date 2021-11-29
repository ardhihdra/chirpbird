package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ardhihdra/chirpbird/chat"
	"github.com/ardhihdra/chirpbird/controllers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var (
	PORT   string
	CLIENT string
	users  = controllers.NewUsersController()
)

func main() {
	loadEnv()
	log.Printf("listening to %s, %s", PORT, CLIENT)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from chirpbird!"))
	})

	mux.HandleFunc("/login", users.Login())

	mux.HandleFunc("/register", users.Register())

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{CLIENT},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}).Handler(mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), handler))

	chat.Start(PORT)
}

func loadEnv() {
	godotenv.Load()
	PORT = os.Getenv("PORT")
	CLIENT = os.Getenv("CLIENT")
	if PORT == "" {
		PORT = "4000"
	}
}
