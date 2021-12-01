package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo/config"
	"github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo/internal/controllers"
	"github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo/internal/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"

	"github.com/joho/godotenv"
)

func main() {
	// Check .env.dev file is existing
	if err := godotenv.Load(".env.dev"); err != nil {
		log.Fatal("failed to load env vars ", err)
	}

	// Create a database connection
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}
	defer config.CloseDatabase(db)

	//init routers
	r := initRoutes(db)

	// Start server
	fmt.Println("Server starting at: 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Server error %v", err)
	}
}

func initRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	dbRepo := repository.NewDBRepo(db)
	friendController := controllers.NewFriendController(dbRepo)

	logger := httplog.NewLogger("friend-management", httplog.Options{
		LogLevel: "trace",
	})
	r.Use(httplog.RequestLogger(logger))

	r.Route("/v1", func(route chi.Router) {
		route.Get("/users", friendController.GetUsers)
		route.Post("/friends", friendController.CreateFriend)
		route.Get("/friends", friendController.GetFriends)
		route.Get("/recipients", friendController.GetRecipientEmails)
		route.Post("/subscription", friendController.CreateSubcription)
		route.Post("/blocking", friendController.CreateUserBlock)
		route.Get("/commonFriends", friendController.GetCommonFriends)
	})

	return r
}
