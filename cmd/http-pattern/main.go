package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/jwtauth/v5"
	"http-pattern/internal/config"
	"http-pattern/internal/http-server/handlers/todo"
	"http-pattern/internal/http-server/middleware/mwlogger"
	"http-pattern/internal/slogger"
	"http-pattern/internal/slogger/sl"
	"http-pattern/internal/storage/psql"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log.Printf("config: %+v", cfg)

	log := slogger.SetupLogger(cfg.Env)
	log.Info("mwlogger started")

	storage, err := psql.New(cfg.Postgres)
	if err != nil {
		slog.Error("cannot create storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mwlogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/login", todo.TODOHandler(log, "/login"))

	log.Info("server started", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("failed to start server", sl.Err(err))
	}

	log.Error("server stopped", sl.Err(err))

	//TODO: HTTP-SERVER
}
