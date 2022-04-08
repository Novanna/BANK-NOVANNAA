package main

import (
	"Trial/BANK-NOVANNA/api"
	"Trial/BANK-NOVANNA/pkg/config"
	"Trial/BANK-NOVANNA/pkg/database"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.GetConfig()
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))

	app := api.SetupRouter(db)
	app.Run(port)
}
