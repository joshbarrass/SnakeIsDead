package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

// Configuration is the env config
type Configuration struct {
	Directory string `envconfig:"DIRECTORY" default:"."`
	Port      int    `envconfig:"PORT" default:"8080"`
}

func main() {
	var config Configuration
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("could not process config: %s", err)
	}
	fmt.Printf("Serving directory %s on port %d.\n", config.Directory, config.Port)
	log.Fatalf("Server exited with err: %s\n", http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port),
		http.FileServer(http.Dir(config.Directory)),
	))
}
