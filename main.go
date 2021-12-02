package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ardhihdra/chirpbird/controllers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var (
	PORT      string
	CLIENT    string
	users     = controllers.NewUsersController()
	groups    = controllers.NewGroupController()
	sessions  = controllers.NewSessionsController()
	messaging = controllers.NewMessagingController()
)

func main() {
	loadEnv()
	log.Printf("listening to %s, %s", PORT, CLIENT)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from chirpbird!"))
	})

	// mux.HandleFunc("/login", users.Login())
	mux.HandleFunc("/register", users.Register())
	mux.HandleFunc("/groups", groups.Create())
	mux.HandleFunc("/groups/:id/join", groups.Join())
	mux.HandleFunc("/groups/:id/left", groups.Left())
	// SESSIONS RESOURCE
	mux.HandleFunc("/sessions", sessions.Create())
	// MESSAGING RESOURCE
	mux.HandleFunc("/:access_token", messaging.Start())

	/** API for Infos */
	mux.HandleFunc("/username", users.CheckUniqueUsername())
	mux.HandleFunc("/dashboard", groups.DashboardData())
	mux.HandleFunc("/rooms", groups.RoomsData())
	mux.HandleFunc("/search", groups.SearchStuff())

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{CLIENT},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}).Handler(mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), handler))
}

func loadEnv() {
	godotenv.Load()
	PORT = os.Getenv("PORT")
	CLIENT = os.Getenv("CLIENT")
	if PORT == "" {
		PORT = "4000"
	}
}
