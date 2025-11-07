package main

import (
	"fmt"
	"inverntory-adi-sanbercode/database"
	"inverntory-adi-sanbercode/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cmd := "start"
	if len(os.Args) < 2 {
		fmt.Println("Gunakan perintah: go run main.go [command]")
		fmt.Println("Command tersedia: migrate | seed | start")
	} else {

		cmd = os.Args[1]
	}

	switch cmd {
	case "migrate":
		database.ConnectDB()
		database.Migrate()
	case "seed":
		database.ConnectDB()
		database.Seed()
	case "start":
		database.ConnectDB()
		routes.SetupRoutes(r)
		r.Run(":8000")
	default:
		log.Fatalf("Command tidak dikenal: %s", cmd)
	}

}
