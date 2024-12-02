package main

import (
	"github.com/ceblay/billing-demo/pkg/ports/http"
	"github.com/ceblay/billing-demo/pkg/service"
)

func main() {
	application := service.NewApplication()
	server := http.NewServer(application)
	server.Run()
}
