package main

import (
	"fmt"
	"log"
	grpc_client "tiny-letter/coordinator/cmd/grpc/client"
	rest_server "tiny-letter/coordinator/cmd/rest/server"
	"tiny-letter/coordinator/pkg/app"
)

func main() {
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%d", config.REST.Domain, config.REST.Port)

	if err := grpc_client.IsGRPC_ClientAvailable(&config.GRPC); err != nil {
		log.Fatal(err.Error())
	}
	rest_server.Serve(addr)
}
