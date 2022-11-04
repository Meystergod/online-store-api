package app

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	_ "github.com/Meystergod/online-store-api/application/docs"
	"github.com/Meystergod/online-store-api/application/internal/config"
	"github.com/Meystergod/online-store-api/application/pkg/logging"
	"github.com/Meystergod/online-store-api/application/pkg/metric"
	"github.com/Meystergod/online-store-api/application/pkg/shutdown"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	cfg        *config.Config
	logger     *logging.Logger
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(cfg *config.Config, logger *logging.Logger) (App, error) {
	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("swagger docs initializing")
	router.Handler(
		http.MethodGet,
		"/swagger",
		http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently),
	)
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logger.Info("heartbeat metric initializing")
	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	return App{
		cfg:    cfg,
		logger: logger,
		router: router,
	}, nil
}

func (s *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return s.startHTTP(ctx)
	})

	return grp.Wait()
}

func (s *App) startHTTP(ctx context.Context) error {
	s.logger.Info("HTTP server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.HTTP.IP, s.cfg.HTTP.Port))
	if err != nil {
		s.logger.WithError(err).Fatal("failed to create listener")
	}

	s.logger.Info("cors initializing")
	c := cors.New(cors.Options{
		AllowedMethods:     s.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins:     s.cfg.HTTP.CORS.AllowedOrigins,
		AllowCredentials:   s.cfg.HTTP.CORS.AllowCredentials,
		AllowedHeaders:     s.cfg.HTTP.CORS.AllowedHeaders,
		OptionsPassthrough: s.cfg.HTTP.CORS.OptionsPassthrough,
		ExposedHeaders:     s.cfg.HTTP.CORS.ExposedHeaders,
		Debug:              s.cfg.HTTP.CORS.Debug,
	})

	handler := c.Handler(s.router)

	s.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	go shutdown.Graceful(s.logger, []os.Signal{
		syscall.SIGABRT,
		syscall.SIGQUIT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		os.Interrupt,
	}, s.httpServer)

	s.logger.Info("application initialized and started")

	if err = s.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			s.logger.Warning("server shutdown")
		default:
			s.logger.Fatal(err)
		}
	}

	return err
}
