package main

import (
	"flag"
	"neon-auth/internal/config"
	"neon-auth/internal/interface/rpc"
	neonrpc "neon-auth/internal/interface/rpc/protocol"
	"neon-auth/internal/register"
	"net"
	"os"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config file")
	flag.Parse()
	server := grpc.NewServer()
	configFile, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	err = configFile.Validate()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	//Build UseCase
	accountUsecase, err := register.AccountUsecaseBuilder(configFile.Database)
	if err != nil {
		logrus.Panic(err)
	}

	//Register GRPC server
	srv := rpc.NewAccountService(accountUsecase)
	neonrpc.RegisterAuthServiceServer(server, srv)
	addr := configFile.RPC.Host + ":" + configFile.RPC.Port

	l, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Service is running on %v", configFile.RPC.Port)
	err = server.Serve(l)
	if err != nil {
		logrus.Fatal(err)
	}
}
