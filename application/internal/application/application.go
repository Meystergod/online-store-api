package app

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	_ "github.com/Meystergod/online-store-api/application/docs"
	"github.com/Meystergod/online-store-api/application/internal/config"
	"github.com/Meystergod/online-store-api/application/internal/controllers/grpc/v1/product"
	"github.com/Meystergod/online-store-api/application/pkg/client/postgresql"
	"github.com/Meystergod/online-store-api/application/pkg/logging"
	"github.com/Meystergod/online-store-api/application/pkg/metric"
	"github.com/Meystergod/online-store-api/application/pkg/shutdown"
	pb_online_store_products "github.com/Meystergod/online-store-grpc-contracts/gen/go/online_store/products/v1"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
)

type App struct {
	cfg                  *config.Config
	router               *httprouter.Router
	httpServer           *http.Server
	grpcServer           *grpc.Server
	pgClient             *pgxpool.Pool
	productServiceServer pb_online_store_products.ProductServiceServer
}

func NewApp(ctx context.Context, cfg *config.Config) (App, error) {
	logger := logging.GetLogger(ctx)

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

	pgConfig := postgresql.NewPgConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port, cfg.PostgreSQL.Database,
	)

	pgClient, err := postgresql.NewClient(ctx, config.PGXPOOL_MAX_ATTEMPTS, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err)
	}

	productServiceServer := product.NewServer(pb_online_store_products.UnimplementedProductServiceServer{})

	return App{
		cfg:                  cfg,
		router:               router,
		pgClient:             pgClient,
		productServiceServer: productServiceServer,
	}, nil
}

func (s *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return s.startHTTP(ctx)
	})
	grp.Go(func() error {
		return s.startGRPC(ctx, s.productServiceServer)
	})
	logging.GetLogger(ctx).Info("application initialized and started")

	return grp.Wait()
}

func (s *App) startGRPC(ctx context.Context, server pb_online_store_products.ProductServiceServer) error {
	logger := logging.GetLogger(ctx).WithFields(map[string]interface{}{
		"IP":   s.cfg.GRPC.IP,
		"PORT": s.cfg.GRPC.Port,
	})
	logger.Info("GRPC server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.GRPC.IP, s.cfg.GRPC.Port))
	if err != nil {
		logger.WithError(err).Fatal("failed to create listener")
	}

	serverOptions := []grpc.ServerOption{}

	s.grpcServer = grpc.NewServer(serverOptions...)

	pb_online_store_products.RegisterProductServiceServer(s.grpcServer, server)

	reflection.Register(s.grpcServer)

	return s.grpcServer.Serve(listener)
}

func (s *App) startHTTP(ctx context.Context) error {
	logger := logging.GetLogger(ctx).WithFields(map[string]interface{}{
		"IP":   s.cfg.HTTP.IP,
		"PORT": s.cfg.HTTP.Port,
	})
	logger.Info("HTTP server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.HTTP.IP, s.cfg.HTTP.Port))
	if err != nil {
		logger.WithError(err).Fatal("failed to create listener")
	}

	logging.GetLogger(ctx).Info("cors initializing")
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

	go shutdown.Graceful(ctx, []os.Signal{
		syscall.SIGABRT,
		syscall.SIGQUIT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		os.Interrupt,
	}, s.httpServer)

	logging.GetLogger(ctx).Info("application initialized and started")

	if err = s.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logging.GetLogger(ctx).Warning("server shutdown")
		default:
			logging.GetLogger(ctx).Fatal(err)
		}
	}

	return err
}
