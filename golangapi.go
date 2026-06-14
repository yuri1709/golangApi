package golangapi

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/yuri1709/golangApi/config"
	"github.com/yuri1709/golangApi/handlers"
)

func init() {
	config.InitFirebase()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handlers.GetUsers)

	functions.HTTP("Api", mux.ServeHTTP)
}
