package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alfiehiscox/submarines/pkg/board"
	"github.com/alfiehiscox/submarines/pkg/html"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
	. "maragu.dev/gomponents"
	ghttp "maragu.dev/gomponents/http"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	if err := start(log); err != nil {
		log.Error("Error starting app:", "error", err)
		os.Exit(1)
	}
}

func start(log *slog.Logger) error {
	log.Info("Starting app")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	server := NewServer(log)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return server.Start()
	})

	<-ctx.Done()
	log.Info("Stopping app")

	eg.Go(func() error {
		return server.Stop()
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	log.Info("App stopped")

	return nil
}

type Server struct {
	log    *slog.Logger
	mux    chi.Router
	server *http.Server
}

func NewServer(log *slog.Logger) *Server {
	mux := chi.NewMux()
	return &Server{
		log: log,
		mux: mux,
		server: &http.Server{
			Addr:              ":8080",
			Handler:           mux,
			ReadTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       5 * time.Second,
		},
	}
}

func (s *Server) setUpRoutes() {
	// Static
	fs := http.FileServer(http.Dir("static"))
	s.mux.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Page Routes
	s.mux.Get("/", ghttp.Adapt(IndexHandler))
	s.mux.Get("/place-ships", ghttp.Adapt(PlaceShipsHandler))
}

func (s *Server) Start() error {
	s.log.Info("Starting HTTP Server", "address", "localhost:8080")
	s.setUpRoutes()
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	s.log.Info("Stopped HTTP Server")
	return nil
}

// Handlers
func PlaceShipsHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return html.PlaceShips(board.NewBoard()), nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return html.Index(), nil
}
