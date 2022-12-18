package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/titoyudha/go_blog_api/api/controllers"
	"github.com/titoyudha/go_blog_api/api/seed"
)

var server = controllers.Server{}

func RunServer() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error gettin' env file : %v", err)
		return
	} else {
		fmt.Println("env loaded")
	}

	server.InitializeDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")

}
