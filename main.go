package main

import (
	"context"
	"fmt"
	"github.com/MruV-RP/mruv-api-placeholder/services"
	"github.com/MruV-RP/mruv-pb-go/accounts"
	"github.com/MruV-RP/mruv-pb-go/business"
	"github.com/MruV-RP/mruv-pb-go/characters"
	"github.com/MruV-RP/mruv-pb-go/devtools"
	"github.com/MruV-RP/mruv-pb-go/economy"
	"github.com/MruV-RP/mruv-pb-go/elevators"
	"github.com/MruV-RP/mruv-pb-go/entrances"
	"github.com/MruV-RP/mruv-pb-go/estates"
	"github.com/MruV-RP/mruv-pb-go/gates"
	"github.com/MruV-RP/mruv-pb-go/groups"
	"github.com/MruV-RP/mruv-pb-go/houses"
	"github.com/MruV-RP/mruv-pb-go/items"
	"github.com/MruV-RP/mruv-pb-go/jobs"
	"github.com/MruV-RP/mruv-pb-go/objects"
	"github.com/MruV-RP/mruv-pb-go/offers"
	"github.com/MruV-RP/mruv-pb-go/organizations"
	"github.com/MruV-RP/mruv-pb-go/plots"
	"github.com/MruV-RP/mruv-pb-go/punishments"
	"github.com/MruV-RP/mruv-pb-go/server"
	"github.com/MruV-RP/mruv-pb-go/spots"
	"github.com/MruV-RP/mruv-pb-go/vehicles"
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
	business.RegisterMruVBusinessServiceServer(s, services.NewBusinessServer())
	characters.RegisterMruVCharactersServiceServer(s, services.NewCharactersServer())
	devtools.RegisterMruVDevToolsServiceServer(s, services.NewDevtoolsServer())
	economy.RegisterMruVEconomyServiceServer(s, services.NewBusinessServer())
	elevators.RegisterMruVElevatorsServiceServer(s, services.NewElevatorsServer())
	entrances.RegisterMruVEntrancesServiceServer(s, services.NewEntrancesServer())
	estates.RegisterMruVEstateServiceServer(s, services.NewEstatesServer())
	gates.RegisterMruVGatesServiceServer(s, services.NewGatesServer())
	groups.RegisterMruVGroupsServiceServer(s, services.NewGroupsServer())
	houses.RegisterMruVHousesServiceServer(s, services.NewHousesServer())
	items.RegisterMruVItemServiceServer(s, services.NewItemsServer())
	jobs.RegisterMruVJobsServiceServer(s, services.NewJobsServer())
	objects.RegisterMruVObjectsServiceServer(s, services.NewObjectsServer())
	offers.RegisterMruVOffersServiceServer(s, services.NewOffersServer())
	organizations.RegisterMruVOrganizationsServiceServer(s, services.NewOrganizationsServer())
	plots.RegisterMruVPlotsServiceServer(s, services.NewPlotsServer())
	punishments.RegisterMruVPunishmentsServiceServer(s, services.NewPunishmentsServer())
	server.RegisterMruVServerServiceServer(s, services.NewServerServer())
	spots.RegisterMruVSpotsServiceServer(s, services.NewSpotsServer())
	vehicles.RegisterMruVVehiclesServiceServer(s, services.NewVehiclesServer())

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
