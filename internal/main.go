package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/vtallen/Basic-Go-REST-API/pkg/swagger/server/restapi"
	"github.com/vtallen/Basic-Go-REST-API/pkg/swagger/server/restapi/operations"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)
	api.GetGopherNameHandler = operations.GetGopherNameHandlerFunc(GetGopherByName)

	// Start listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}

func GetGopherByName(gopher operations.GetGopherNameParams) middleware.Responder {
	var URL string
	if gopher.Name != "" {
		URL = "https://raw.githubusercontent.com/ashleymcnamara/gophers/master/" + gopher.Name + ".png"
		fmt.Println(URL)
	} else {
		URL = "https://github.com/ashleymcnamara/gophers/blob/master/DEATH_METAL_GOPHER.png"
	}

	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("error")
	}

	return operations.NewGetGopherNameOK().WithPayload(response.Body)
}
