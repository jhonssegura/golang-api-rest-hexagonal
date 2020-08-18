package di

import (
	"golang-api-rest-hexagonal/pkg/config"
	"golang-api-rest-hexagonal/pkg/storage"

	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)
	// DB
	container.Provide(storage.NewDb)
	// Ejemplo
	//container.Provide(orm.NewServerRepo)
	//container.Provide(server.NewServerService)
	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
