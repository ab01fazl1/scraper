package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ab01fazl1/scraper/internal/database"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries	
}

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("env file not found")
	}

	// database conn
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Println("DB_URL not found in env file")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("can't connect to db", err)
	}
	
	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	portaddr := os.Getenv("PORT")
	if portaddr == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()	// I did not include the cors headers	FYI
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUserByApiKey)
	
	server := &http.Server{
		Addr: ":" + portaddr,
		Handler: router,
	}


	log.Printf("Server running on port: %v" , portaddr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	

}
