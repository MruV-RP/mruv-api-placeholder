package main

import (
	"context"
	"fmt"
	"github.com/MruV-RP/mruv-api-placeholder/services"
	"github.com/MruV-RP/mruv-pb-go/accounts"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func RunGRPCServer() {
	// if we crash the go code, we get the file name and line number in log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logrus.Infoln("Starting listener for gRPC API")
	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", viper.GetString("HOST"), viper.GetInt("GRPC_PORT")))
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}
	logrus.Infoln("Listener started on", listener.Addr())

	//serve gRPC services
	s := grpc.NewServer(grpc.MaxRecvMsgSize(1048576))
	reflection.Register(s)
	defer s.Stop()

	//Register services
	accounts.RegisterMruVAccountsServiceServer(s, services.NewAccountsServer())

	go func() {
		logrus.Println("Starting server.")

		if err := s.Serve(listener); err != nil {
			logrus.Fatalln("Failed to serve", err)
		}
	}()

	// Set up the gRPC gateway for REST endpoints
	if err = setUpgRPCGateway(); err != nil {
		logrus.Fatalln("Failed to set up gRPC gateway: ", err)
	}

	// Wait for CTRL+C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch // Block until signal is received
	logrus.Infoln("\nStopping the server.")
}

func setUpgRPCGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()
	err := accounts.RegisterMruVAccountsServiceHandlerFromEndpoint(ctx, mux,
		fmt.Sprintf("%s:%d", viper.GetString("HOST"), viper.GetInt("GRPC_PORT")),
		opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", viper.GetString("HOST"), viper.GetInt("HTTP_PORT")), mux)
}

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("GRPC_PORT", 8081)

	RunGRPCServer()
}
