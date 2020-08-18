package main

import (
	"fmt"
	"os"

	"golang-api-rest-hexagonal/cmd/appserver"
	"golang-api-rest-hexagonal/pkg/config"
	"golang-api-rest-hexagonal/pkg/di"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {

	g := gin.Default()
	d := di.BuildContainer()
	c, _ := config.NewConfig()

	g.Use(cors.Default())
	svr := appserver.NewAppServer(g, d)
	svr.MapRoutes(c)
	if err := svr.SetupDB(); err != nil {
		return err
	}
	return svr.Start()
}
