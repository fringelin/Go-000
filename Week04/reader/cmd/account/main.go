package main

import (
	"context"
	"log"
	"reader/api/account"
	"reader/internal/account/di"

	"github.com/go-kratos/kratos/v2"
	_ "github.com/go-kratos/kratos/v2/encoding/json"
	_ "github.com/go-kratos/kratos/v2/encoding/proto"
	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string
	// Branch is current branch name the code is built off.
	Branch string
	// Revision is the short commit hash of source tree.
	Revision string
	// BuildDate is the date when the binary was built.
	BuildDate string
)

func main() {
	log.Printf("service version: %s\n", Version)

	// transport server
	grpcSrv := grpctransport.NewServer(":9001")

	// register service
	gs, cleanup, err := di.NewAccountService()
	if err != nil {
		panic(err)
	}
	account.RegisterAccountServer(grpcSrv, gs)

	// application lifecycle
	app := kratos.New()
	app.Append(kratos.Hook{OnStart: nil, OnStop: func(_ context.Context) error {
		cleanup()
		return nil
	}})
	app.Append(kratos.Hook{OnStart: grpcSrv.Start, OnStop: grpcSrv.Stop})

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Printf("startup failed: %v\n", err)
	}
}
