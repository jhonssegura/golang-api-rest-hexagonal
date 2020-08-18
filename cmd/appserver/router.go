package appserver

import (
	"fmt"
	"net/http"

	"golang-api-rest-hexagonal/pkg/config"
	"golang-api-rest-hexagonal/pkg/http/rest"

	"github.com/gin-gonic/gin"
)

func (as *appserver) MapRoutes(c *config.Config) {

	// Group : v1
	apiV1 := as.router.Group("rest-sgc/api/v1")

	as.healthRoutes(apiV1)
	as.metricRoutes(apiV1, c)
	//as.serverRoutes(apiV1)
}

func (as *appserver) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")
	{
		h := rest.NewHealthCtrl()
		healthRoutes.GET("/", h.Ping)
	}
}

func (as *appserver) metricRoutes(api *gin.RouterGroup, d *config.Config) {
	metricRoutes := api.Group("/metrics")
	{
		metricRoutes.GET("/ict", func(c *gin.Context) {
			c.String(http.StatusOK, "{\"indicator\":\"ICT\",\"value\":0.503,\"drivers\":[{\"name\":\"Volumen\",\"weight\":0.45,\"kpi\":0.271429,\"complexityVariables\":[{\"variable\":\"Número de servidores gestionados - No SPSA\",\"weight\":0.5,\"targetValue\":35,\"currentValue\":53,\"kpi\":0.51428574},{\"variable\":\"Número de servidores SPSA compartido con otras empresas\",\"weight\":0.5,\"targetValue\":35,\"currentValue\":54,\"kpi\":0.54285717}]},{\"name\":\"Diversidad\",\"weight\":\"0.45\",\"kpi\":\"0.624083\",\"complexityVariables\":[{\"variable\":\"Número de aplicaciones no estandarizadas\",\"weight\":0.05,\"targetValue\":20,\"currentValue\":41,\"kpi\":1.05},{\"variable\":\"Número de aplicaciones con ciclos de vida no automatizados\",\"weight\":0.1,\"targetValue\":30,\"currentValue\":42,\"kpi\":0.4},{\"variable\":\"Número de aplicaciones sin administración de acceso centralizada\",\"weight\":0.05,\"targetValue\":30,\"currentValue\":42,\"kpi\":0.4},{\"variable\":\"Número de aplicaciones sin administración de permisos centralizada\",\"weight\":0.05,\"targetValue\":30,\"currentValue\":42,\"kpi\":0.4},{\"variable\":\"Número de aplicaciones sin código fuente centralizado\",\"weight\":0.15,\"targetValue\":30,\"currentValue\":42,\"kpi\":0.4},{\"variable\":\"Número de flujos programados fuera del estándar \",\"weight\":0.05,\"targetValue\":100,\"currentValue\":169,\"kpi\":0.69},{\"variable\":\"Número de Servidores con SO con versiones obsoletas\",\"weight\":0.2,\"targetValue\":18,\"currentValue\":34,\"kpi\":0.8888889},{\"variable\":\"Número de Servidores no virtualizados\",\"weight\":0.05,\"targetValue\":30,\"currentValue\":50,\"kpi\":0.6666667},{\"variable\":\"Número de Servidores con componentes de infraestructura no estándar\",\"weight\":0.05,\"targetValue\":40,\"currentValue\":51,\"kpi\":0.275},{\"variable\":\"Número de Bases de Datos con versiones obsoletas o no soportadas\",\"weight\":0.2,\"targetValue\":27,\"currentValue\":42,\"kpi\":0.5555556},{\"variable\":\"Número de Bases de Datos no estandarizadas\",\"weight\":0.05,\"targetValue\":9,\"currentValue\":20,\"kpi\":1.2222222}]},{\"name\":\"Interdependencia\",\"weight\":\"0.1\",\"kpi\":\"1\",\"complexityVariables\":[{\"variable\":\"Número de aplicaciones no adaptables a la nube\",\"weight\":1,\"targetValue\":24,\"currentValue\":48,\"kpi\":1}]}]}")
		})

		metricRoutes.GET("/temp", func(c *gin.Context) {
			tast := fmt.Sprintf("app=%s", d.AppName)
			c.String(http.StatusOK, tast)
		})
	}
}

// Ejemplo
/*
func (as *appserver) serverRoutes(api *gin.RouterGroup) {
	serverRoutes := api.Group("/servers")
	{
		var serverSvc server.Service
		var dbSvc database.Service
		as.cont.Invoke(func(s server.Service, d database.Service) {
			serverSvc = s
			dbSvc = d
		})

		svr := rest.NewServerCtrl(serverSvc, dbSvc)

		serverRoutes.GET("/:id", svr.GetByID)
		serverRoutes.GET("/:id/databases", svr.GetDBByID)
		serverRoutes.GET("/", svr.GetAll)
		serverRoutes.POST("/", svr.Store)
		serverRoutes.PUT("/:id", svr.Update)
		serverRoutes.DELETE("/:id", svr.Delete)

	}
}
*/
