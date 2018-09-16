//go:generate go run -tags=dev assets_generate.go

package main

import (
	"fmt"
	"net/http"

	"github.com/jsorrell/www.jacksorrell.com/configloader"
	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/jsorrell/www.jacksorrell.com/web"
)

func main() {
	server := web.NewServer()

	log.WithField("port", configloader.Config.Server.Port).Info("Starting server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configloader.Config.Server.Port), server))
}
