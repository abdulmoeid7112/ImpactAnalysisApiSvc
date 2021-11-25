package main

import (
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	runtime "github.com/abdulmoeid7112/impact-analysis-api-svc"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/config"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/gen/restapi"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(context.TODO(), rt, swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	defer server.Shutdown()


	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
