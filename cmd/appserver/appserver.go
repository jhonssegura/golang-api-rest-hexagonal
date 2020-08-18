package appserver

import (
	"fmt"
	"golang-api-rest-hexagonal/pkg/config"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type appserver struct {
	router *gin.Engine
	cont   *dig.Container
}

func NewAppServer(e *gin.Engine, c *dig.Container) *appserver {
	return &appserver{
		router: e,
		cont:   c,
	}
}

func (as *appserver) SetupDB() error {
	var db *gorm.DB

	if err := as.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}

	db.AutoMigrate()

	return nil
}

// Start serving the application
func (as *appserver) Start() error {
	var cfg *config.Config
	if err := as.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return as.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
