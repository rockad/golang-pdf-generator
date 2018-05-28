package app

import (
	"fmt"
	"golang-pdf-generator/app/routes"
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

var Server *echo.Echo

func readConfig() {
	viper.SetConfigType("json")
	viper.SetConfigName("app-config")
	viper.AddConfigPath(path.Join(".", "app", "config"))

	err := viper.ReadInConfig()
	if err != nil {
		Server.Logger.Fatalf("Fatal error config file: %s \n", err)
	}
}

// Init initialize application
func Init() {
	appConf := viper.GetStringMapString("app")

	// Middleware
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())

	routes.Init(Server)

	Server.Logger.Fatal(Server.Start(fmt.Sprintf("%s:%s", appConf["host"], appConf["port"])))
}

func init() {
	Server = echo.New()
	readConfig()
}
