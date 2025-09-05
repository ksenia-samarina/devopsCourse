package main

import (
	"context"
	"devopsCourse/internal/config"
	"devopsCourse/internal/domains/order"
	"devopsCourse/internal/handlers"
	"devopsCourse/internal/repository/postgres"
	"devopsCourse/internal/repository/postgres/transactor"
	"devopsCourse/libs/interceptors"
	lomsv1 "devopsCourse/pkg/loms_v1"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("error config init: %s", err.Error())
	}

	// transport (gRPC listener)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.PortGRPC))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	// Postgres pool
	pool, err := pgxpool.Connect(ctx, cfg.Storages.Postgres)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer pool.Close()

	maxConnLifetime, err := time.ParseDuration(cfg.Storages.MaxConnLifetime)
	if err != nil {
		log.Fatalf("invalid format maxConnLifetime")
	}

	maxConnIdleTime, err := time.ParseDuration(cfg.Storages.MaxConnLifetime)
	if err != nil {
		log.Fatalf("invalid format maxConnIdleTime")
	}

	poolCfg := pool.Config()
	poolCfg.MinConns = cfg.Storages.MinConnections
	poolCfg.MaxConns = cfg.Storages.MaxConnections
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime

	// transactor and repository
	transManager := transactor.New(pool)
	storage := postgres.New(transManager)

	// domains
	orderDomain := order.New(storage, transManager)

	// gRPC server
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors.LogServerInterceptor))

	// register handlers
	reflection.Register(grpcServer)
	lomsv1.RegisterLomsV1Server(grpcServer, handlers.New(orderDomain))

	// запускаем gRPC сервер в отдельной горутине
	go func() {
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// REST Gateway
	mux := runtime.NewServeMux()
	grpcAddr := fmt.Sprintf(":%s", cfg.PortGRPC)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := lomsv1.RegisterLomsV1HandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	restAddr := cfg.PortREST
	log.Printf("REST gateway listening at %s", restAddr)
	if err := http.ListenAndServe(restAddr, mux); err != nil {
		log.Fatalf("failed to serve REST gateway: %v", err)
	}
}
