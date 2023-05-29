package main

import (
	"example.com/go-chat/database"
	"example.com/go-chat/internal/WS"
	"example.com/go-chat/internal/user"
	"example.com/go-chat/router"
	"log"
)

func main() {
	conn, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Cannot initialize database", err)
	}
	userRep := user.NewRepository(conn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := WS.NewHub()
	wshandler := WS.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wshandler)
	router.Start("localhost:8080")
}
