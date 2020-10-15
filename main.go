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
		service.Version("latest"),
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
