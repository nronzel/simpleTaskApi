package main

// @title Simple Task API
// @version 1.0
// @description Task/To-Do management API. This API allows you to create, read, update, and delete tasks.
// @host localhost:9090
// @BasePath /

import (
	"net/http"
	"os"
	_ "simpleTaskApi/docs"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func main() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(consoleWriter).With().Timestamp().Logger()

	db := tempDB{
		tasks:  make(map[int]taskEntry),
		nextID: 1,
		logger: logger,
	}
	mux := chi.NewRouter()

	mux.Use(db.logRequestMiddleware)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to the Simple Task API!\nCheck out the API documentation at localhost:9090/swagger/\n"))
		if err != nil {
			logger.Error().Err(err).Msg("Error writing to response")
		}
	})

	// Tasks
	mux.Get("/tasks", db.handlerTasksGetAll)
	mux.Get("/tasks/{taskID}", db.handlerTasksGet)
	mux.Post("/tasks", db.handlerTasksPost)
	mux.Delete("/tasks/{taskID}", db.handlerTasksDelete)
	mux.Put("/tasks/{taskID}", db.handlerTasksUpdate)

	// redirect to swagger
	mux.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	mux.Get("/swagger/*", httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler:      mux,
		Addr:         ":9090",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	logger.Info().Msg("Server listening on localhost:9090...")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server stopped unexpectedly!")
	}
}
