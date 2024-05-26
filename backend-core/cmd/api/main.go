package main

import (
	"backend-core/internal/server"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	server := server.New()
	server.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
