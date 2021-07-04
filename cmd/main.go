package main

import (
	"log"

	todo "github.com/DT8/app_todo"
	"github.com/DT8/app_todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
