package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-http/service"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "spices.dev/stollenaar/http",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), service.New("0.0.3"), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
