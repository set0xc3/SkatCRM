package server

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Context struct {
	Host string
	Port string
}

func New() (ctx Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx = Context{
		Host: os.Getenv("SERVER_HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}

	if !strings.HasPrefix(ctx.Port, ":") {
		ctx.Port = ":" + ctx.Port
	}
	return
}
