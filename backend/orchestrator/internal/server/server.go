package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"orchstrator/internal/server/handlers"
	mdb "orchstrator/pkg/db/mongodb"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Server struct {
	E       *echo.Echo
	MongoDB *mdb.Mongo

	// Context *context.Context
}

func New(MongoDB *mdb.Mongo) *Server {
	e := echo.New()
	return &Server{E: e, MongoDB: MongoDB}
}

func (s *Server) Start(port string) {

	log.Info("starting on ", port)

	s.E.Logger.SetLevel(log.INFO)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	s.applyMiddlewares()
	s.applyHandlers()

	go func() {
		if err := s.E.Start(port); err != nil && err != http.ErrServerClosed {
			log.Error("server error : ", err)
			s.E.Logger.Fatal("shutting down the server ...")
		}
	}()

	<-ctx.Done()
}

func (s *Server) GracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.E.Shutdown(ctx); err != nil {
		s.E.Logger.Fatal(err)
	}
}

func (s *Server) applyMiddlewares() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	s.E.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	s.E.Use(middleware.Recover())
}

func (s *Server) applyHandlers() {
	h := handlers.NewRouter(s.MongoDB)

	s.E.GET("api/v1/ping", h.GetPing)
}
