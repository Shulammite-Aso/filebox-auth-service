package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Shulammite-Aso/filebox-auth-service/pkg/config"
	"github.com/Shulammite-Aso/filebox-auth-service/pkg/db"
	"github.com/Shulammite-Aso/filebox-auth-service/pkg/proto"
	"github.com/Shulammite-Aso/filebox-auth-service/pkg/service"
	"github.com/Shulammite-Aso/filebox-auth-service/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "filebox-auth-service",
		ExpirationHours: 24 * 1,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	fmt.Println("Auth service listening on", c.Port)

	s := service.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
