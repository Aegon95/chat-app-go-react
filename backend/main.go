package main

import (
	_ "github.com/Aegon95/chat-app-go-react/docs"
	"github.com/Aegon95/chat-app-go-react/internal/api"
)

// @Golang API REST
// @version 1.0
// @description API REST in Golang with Gin Framework

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	api.Run("")
}
