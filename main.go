//go:generate protoc -Iproto --micro_out=paths=source_relative:./proto --go_out=paths=source_relative:./proto proto/persona/realm.proto
//go:generate protoc -Iproto --micro_out=paths=source_relative:./proto --go_out=paths=source_relative:./proto proto/persona/group.proto
//go:generate protoc -Iproto --micro_out=paths=source_relative:./proto --go_out=paths=source_relative:./proto proto/persona/role.proto
//go:generate protoc -Iproto --micro_out=paths=source_relative:./proto --go_out=paths=source_relative:./proto proto/persona/user.proto
package main

import (
	"context"
	"os"
	"time"

	"github.com/johnbellone/persona-service/internal/handler"
	pb "github.com/johnbellone/persona-service/proto/persona"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/micro/trace/uuid"
	"github.com/micro/micro/plugin"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Tracef("[persona] request: %v", req.Endpoint())
		err := fn(ctx, req, rsp)
		return err
	}
}

func init() {
	if err := plugin.Register(uuid.New()); err != nil {
		panic(err)
	}
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.persona"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(logWrapper),
	)

	service.Init()
	if err := pb.RegisterUserServiceHandler(service.Server(), new(handler.UserService)); err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
		os.Exit(3)
	}
}
