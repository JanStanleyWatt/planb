package main

import (
	"net/http"

	"github.com/JanStanleyWatt/planb/pkg/autogen/go/api/v1/apiv1connect"
	"github.com/JanStanleyWatt/planb/server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	srv := &server.PlanBServiceServer{}
	mux := http.NewServeMux()

	path, handler := apiv1connect.NewPlanBClientToServiceHandler(srv)
	mux.Handle(path, handler)

	http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
}
