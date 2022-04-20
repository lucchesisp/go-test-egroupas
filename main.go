package main

import (
	"fmt"
	"github.com/lucchesisp/go-test-egroupas/src/config"
	"github.com/lucchesisp/go-test-egroupas/src/routes"
)

func main() {
	port := config.GetEnvVariable("PORT")

	fmt.Printf("Starting server on port %s\n", port)
	routes.Run(port)
}
