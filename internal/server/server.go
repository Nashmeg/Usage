package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/Nashmeg/Usage.git/internal/api"
	"github.com/gorilla/mux"
)

// Start starts the server.  It does not return unless there is a failure.
func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("'PORT' environment variable not set. This is the port where this program will be available.")
	}

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: Router(),
	}
	log.Println("System running on localhost:" + port + "/api/usage/playground")
	log.Fatal(http.ListenAndServe(":"+port, Router()))

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Printf("Failed starting or closing server.  Error: %v", err)
	}
}

// Router manages all the URL routes.
func Router() *mux.Router {
	router := mux.NewRouter()
	// publicPaths := []string{
	// 	"/api/usage/health",
	// }
	// router.Use(session.Middleware(publicPaths))
	router.HandleFunc("/api/usage/playground", handler.Playground("GraphQL playground", "/api/usage"))
	router.HandleFunc("/api/usage", handler.GraphQL(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{}})))
	router.HandleFunc("/api/usage/health", getHealth)
	return router
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
