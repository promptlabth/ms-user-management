package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/promptlabth/ms-user-management/app/user"
	"github.com/promptlabth/ms-user-management/config"
	"github.com/promptlabth/ms-user-management/logger"
	userProto "github.com/promptlabth/proto-lib/user"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()

	logger.InitLogger()

	tp, err := config.InitTrace()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	grpcEndpoint := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// create a grpc server
	grpcServe := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	gormDB, err := initDb()
	if err != nil {
		logger.Fatal(ctx, err.Error())
	}

	userRepo := user.NewRepostiroy(gormDB)
	userService := user.NewService(userRepo)
	userRegis := user.NewRegister(userService)

	userProto.RegisterUserServiceServer(grpcServe, userRegis)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		<-ctx.Done()
		grpcServe.GracefulStop()
		defer wg.Done()
	}()

	log.Printf("starting grpc server at endpoint %s", grpcEndpoint)
	err = grpcServe.Serve(lis)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
	wg.Wait()

}

func initDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
