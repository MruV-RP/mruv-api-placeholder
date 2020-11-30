package main

import (
	"context"
	"fmt"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
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
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
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

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("GRPC_PORT", 8081)

	RunGRPCServer()
}

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
	s := grpc.NewServer(grpc.MaxRecvMsgSize(1048576),
		grpc.StreamInterceptor(middleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(logrus.NewEntry(logrus.New())),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logrus.New())),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	reflection.Register(s)
	defer s.Stop()

	//Register services
	gen := &generator.SimpleGenerator{}
	accounts.RegisterMruVAccountsServiceServer(s, services.NewAccountsServer(gen))
	business.RegisterMruVBusinessServiceServer(s, services.NewBusinessServer(gen))
	characters.RegisterMruVCharactersServiceServer(s, services.NewCharactersServer(gen))
	devtools.RegisterMruVDevToolsServiceServer(s, services.NewDevtoolsServer(gen))
	economy.RegisterMruVEconomyServiceServer(s, services.NewEconomyServer(gen))
	elevators.RegisterMruVElevatorsServiceServer(s, services.NewElevatorsServer(gen))
	entrances.RegisterMruVEntrancesServiceServer(s, services.NewEntrancesServer(gen))
	estates.RegisterMruVEstateServiceServer(s, services.NewEstatesServer(gen))
	gates.RegisterMruVGatesServiceServer(s, services.NewGatesServer(gen))
	groups.RegisterMruVGroupsServiceServer(s, services.NewGroupsServer(gen))
	houses.RegisterMruVHousesServiceServer(s, services.NewHousesServer(gen))
	items.RegisterMruVItemServiceServer(s, services.NewItemsServer(gen))
	jobs.RegisterMruVJobsServiceServer(s, services.NewJobsServer(gen))
	objects.RegisterMruVObjectsServiceServer(s, services.NewObjectsServer(gen))
	offers.RegisterMruVOffersServiceServer(s, services.NewOffersServer(gen))
	organizations.RegisterMruVOrganizationsServiceServer(s, services.NewOrganizationsServer(gen))
	plots.RegisterMruVPlotsServiceServer(s, services.NewPlotsServer(gen))
	punishments.RegisterMruVPunishmentsServiceServer(s, services.NewPunishmentsServer(gen))
	server.RegisterMruVServerServiceServer(s, services.NewServerServer(gen))
	spots.RegisterMruVSpotsServiceServer(s, services.NewSpotsServer(gen))
	vehicles.RegisterMruVVehiclesServiceServer(s, services.NewVehiclesServer(gen))

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
	eptString := fmt.Sprintf("%s:%d", viper.GetString("HOST"), viper.GetInt("GRPC_PORT"))
	if err := accounts.RegisterMruVAccountsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := business.RegisterMruVBusinessServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := characters.RegisterMruVCharactersServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := economy.RegisterMruVEconomyServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := elevators.RegisterMruVElevatorsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := entrances.RegisterMruVEntrancesServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := estates.RegisterMruVEstateServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := gates.RegisterMruVGatesServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := groups.RegisterMruVGroupsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := houses.RegisterMruVHousesServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := items.RegisterMruVItemServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := jobs.RegisterMruVJobsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := objects.RegisterMruVMovableObjectsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := objects.RegisterMruVObjectModelsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := objects.RegisterMruVObjectsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := offers.RegisterMruVOffersServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := organizations.RegisterMruVOrganizationsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := plots.RegisterMruVPlotsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := punishments.RegisterMruVPunishmentsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := server.RegisterMruVServerServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := spots.RegisterMruVSpotsServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}
	if err := vehicles.RegisterMruVVehiclesServiceHandlerFromEndpoint(ctx, mux, eptString, opts); err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", viper.GetString("HOST"), viper.GetInt("HTTP_PORT")), mux)
}
