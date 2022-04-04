package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ab3llo/go-auth/pkg/auth/pb"
	"github.com/ab3llo/go-auth/pkg/config"
	"github.com/ab3llo/go-auth/pkg/db"
	"github.com/ab3llo/go-auth/pkg/services"
	"github.com/ab3llo/go-auth/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}
	database := db.Connect(&cfg)
	jwt := utils.JwtWrapper{
		SecretKey:       cfg.JwtSecretKey,
		Issuer:          "go-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", cfg.Port)

	if err != nil {
		log.Fatalln("Server failed to listen:", err)
	}

	fmt.Println("Auth svc on", cfg.Port)

	s := services.Server{
		DbConnection: database,
		Jwt:          jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to start server: ", err)
	}
}
