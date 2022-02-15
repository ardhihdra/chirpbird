package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ardhihdra/chirpbird/app/controllers"
	"github.com/ardhihdra/chirpbird/app/messaging"
	"github.com/ardhihdra/chirpbird/app/models"
	"github.com/ardhihdra/chirpbird/app/repository"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var (
	PORT        string
	CLIENT      string
	groupRepo   = repository.NewGroupElasticRepository()
	userRepo    = repository.NewUserElasticRepository()
	sessionRepo = repository.NewSessionElasticRepository()
	messageRepo = repository.NewMessageElasticRepository()

	usersModel   = models.NewUsersHandler(userRepo)
	groupModel   = models.NewGroupsModel(groupRepo)
	sessionModel = models.NewSessionsModel(sessionRepo)
	messageModel = models.NewMessageModel(messageRepo)

	users    = controllers.NewUsersController(usersModel)
	groups   = controllers.NewGroupController(groupModel)
	sessions = controllers.NewSessionsController(sessionModel)
)

func main() {
	loadEnv()
	log.Printf("listening to %s, %s", PORT, CLIENT)

	mux := http.NewServeMux()

	initRoute(mux)

	// allowedHeaders := []string{"Accept", "User-Agent", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token", "Origin"}
	allowedHeaders := []string{"Accept", "Content-Type", "Origin", "*"}
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{CLIENT},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   allowedHeaders,
		AllowCredentials: true,
	}).Handler(mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), handler))

	messaging.NewMessagingService(usersModel, groupModel, sessionModel, messageModel)
}

func loadEnv() {
	godotenv.Load()
	PORT = os.Getenv("PORT")
	CLIENT = os.Getenv("CLIENT")
	if PORT == "" {
		PORT = "4000"
	}
	if CLIENT == "" {
		CLIENT = "*"
	}
}

func initRoute(mux *http.ServeMux) {
	// mux.HandleFunc("/login", users.Login())
	mux.HandleFunc("/register", users.Register())
	mux.HandleFunc("/logout", users.Logout())
	mux.HandleFunc("/groups", groups.Create())
	mux.HandleFunc("/groups/:id/join", groups.Join())
	mux.HandleFunc("/groups/:id/left", groups.Left())

	/** API for Infos */
	mux.HandleFunc("/users", users.GetUsers())
	mux.HandleFunc("/username", users.CheckUniqueUsername())
	mux.HandleFunc("/dashboard", groups.DashboardData())
	mux.HandleFunc("/rooms", groups.RoomsData())
	mux.HandleFunc("/search", groups.SearchStuff())

	// SESSIONS RESOURCE
	mux.HandleFunc("/sessions", sessions.Create())
	// MESSAGING RESOURCE
	mux.HandleFunc("/messaging", messaging.Start())
}
