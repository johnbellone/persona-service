//go:generate protoc --experimental_allow_proto3_optional -Iproto --go_out=paths=source_relative:./proto proto/persona/type/realm.proto proto/persona/type/group.proto proto/persona/type/role.proto proto/persona/type/user.proto proto/persona/type/person.proto
//go:generate protoc -Iproto --micro_out=paths=source_relative:./proto --go_out=paths=source_relative:./proto proto/persona/api/auth.proto proto/persona/api/realm.proto proto/persona/api/group.proto proto/persona/api/role.proto proto/persona/api/user.proto
package main

import (
	"context"
	"os"
	"time"

	"github.com/johnbellone/persona-service/internal/handler"
	pb "github.com/johnbellone/persona-service/proto/persona/api"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/micro/trace/uuid"
	"github.com/micro/micro/plugin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	// TODO: Add support for configuring error capture and tracing with Sentry.

	_, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		os.Exit(2)
	}

	if err := pb.RegisterGroupServiceHandler(service.Server(), new(handler.GroupService)); err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	if err := pb.RegisterRoleServiceHandler(service.Server(), new(handler.RoleService)); err != nil {
		log.Fatal(err)
		os.Exit(3)
	}

	if err := pb.RegisterRealmServiceHandler(service.Server(), new(handler.RealmService)); err != nil {
		log.Fatal(err)
		os.Exit(4)
	}

	if err := pb.RegisterUserServiceHandler(service.Server(), new(handler.UserService)); err != nil {
		log.Fatal(err)
		os.Exit(5)
	}

	if err := pb.RegisterAuthServiceHandler(service.Server(), new(handler.AuthService)); err != nil {
		log.Fatal(err)
		os.Exit(6)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
