package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/oauth2/clientcredentials"

	"bitbucket.org/phoops/vilma/internal/core/interactors"
	"bitbucket.org/phoops/vilma/internal/infrastructure/keycloak"
	"bitbucket.org/phoops/vilma/internal/infrastructure/log"
	"bitbucket.org/phoops/vilma/internal/infrastructure/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_server "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"bitbucket.org/phoops/vilma/internal/infrastructure/config"
	"bitbucket.org/phoops/vilma/internal/infrastructure/gateways/grpc"
)

var Version = "development"
var BuildDate = time.Now().Format("Mon Jan 2 15:04:05")

const splashScreen = `
                                            ___           ___
     ___        ___                        /__/\         /  /\    
    /__/\      /  /\                      |  |::\       /  /::\   
    \  \:\    /  /:/      ___     ___     |  |:|:\     /  /:/\:\  
     \  \:\  /__/::\     /__/\   /  /\  __|__|:|\:\   /  /:/~/::\ 
 ___  \__\:\ \__\/\:\__  \  \:\ /  /:/ /__/::::| \:\ /__/:/ /:/\:\
/__/\ |  |:|    \  \:\/\  \  \:\  /:/  \  \:\~~\__\/ \  \:\/:/__\/
\  \:\|  |:|     \__\::/   \  \:\/:/    \  \:\        \  \::/     
 \  \:\__|:|     /__/:/     \  \::/      \  \:\        \  \:\     
  \__\::::/      \__\/       \__\/        \  \:\        \  \:\    
      ~~~~                                 \__\/         \__\/    

			Your nordic user info retriever. 
			Version: %s BuildDate: %s
`

func main() {
	config, err := config.LoadVilmaConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf(splashScreen, Version, BuildDate)

	logger, err := log.NewLogger(config.Environment == "production", "vilma")
	if err != nil {
		panic(err)
	}

	logger.Info("starting vilma")

	credentialsConfig := clientcredentials.Config{
		ClientID:     config.OAuth2ClientID,
		ClientSecret: config.OAuth2ClientSecret,
		TokenURL:     config.OAuth2TokenURL,
	}

	keycloakService := keycloak.NewService(
		logger,
		credentialsConfig.TokenURL,
		credentialsConfig.ClientID,
		credentialsConfig.ClientSecret,
		config.KeycloakRealm,
		config.KeycloakBaseUrl,
	)

	getIdentityByIdUsecase := interactors.NewGetIdentityById(logger, keycloakService)

	grpcHandlers := grpc.NewVilmaHandler(logger, getIdentityByIdUsecase)

	lis, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc_server.NewServer(
		grpc_server.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger.Desugar()),
		)),
	)
	proto.RegisterVilmaIdentityPoolServer(s, grpcHandlers)
	if config.Environment != "production" {
		reflection.Register(s)
	}
	// run signal handling goroutine
	go grpc.GracefulShutdown(logger, s)

	logger.Infof("listening on %s", config.ListenAddr)
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
