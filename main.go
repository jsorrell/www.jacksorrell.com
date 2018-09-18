//go:generate go run -tags=dev assets_generate.go

package main

import (
	"fmt"
	"net/http"

	"github.com/jsorrell/www.jacksorrell.com/config"
	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/jsorrell/www.jacksorrell.com/web"
)

func main() {
	server := web.NewServer()

	log.WithField("port", config.Server.Port).Info("Starting server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), server))
}
