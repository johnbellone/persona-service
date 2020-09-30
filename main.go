//go:generate protoc --experimental_allow_proto3_optional -Iproto --go_out=paths=source_relative:./proto proto/persona/type/realm.proto proto/persona/type/group.proto proto/persona/type/role.proto proto/persona/type/user.proto proto/persona/type/person.proto
//go:generate protoc -Iproto --micro_out=paths=source_relative:./proto --go_out=paths=source_relative:./proto proto/persona/api/auth.proto proto/persona/api/realm.proto proto/persona/api/group.proto proto/persona/api/role.proto proto/persona/api/user.proto
package main

import (
	"os"

	"github.com/johnbellone/persona-service/internal/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	srv := service.New(
		service.Name("persona"),
	)

	_, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		os.Exit(2)
	}

	if err := srv.Handle(new(handler.AuthService)); err != nil {
		logger.Fatal(err)
		os.Exit(3)
	}

	if err := srv.Handle(new(handler.GroupService)); err != nil {
		logger.Fatal(err)
		os.Exit(4)
	}

	if err := srv.Handle(new(handler.RoleService)); err != nil {
		logger.Fatal(err)
		os.Exit(5)
	}

	if err := srv.Handle(new(handler.RealmService)); err != nil {
		logger.Fatal(err)
		os.Exit(6)
	}

	if err := srv.Handle(new(handler.UserService)); err != nil {
		logger.Fatal(err)
		os.Exit(7)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}
