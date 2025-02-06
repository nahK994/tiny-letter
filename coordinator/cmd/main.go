package main

import (
	"log"
	grpc_client "tiny-letter/coordinator/cmd/grpc/client"
	rest_server "tiny-letter/coordinator/cmd/rest/server"
	"tiny-letter/coordinator/pkg/app"
)

func main() {
	config := app.GetConfig()

	if err := grpc_client.IsGRPC_ClientAvailable(&config.GRPC); err != nil {
		log.Fatal(err.Error())
	}
	rest_server.Serve(&config.REST)
}
