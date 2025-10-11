package internal

import (
	"net/http"

	"cloud.google.com/go/spanner"
)

func NewHandler(client *spanner.Client) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create-user", CreateUserHandler(client))
	mux.HandleFunc("/get-user", GetUserHandler(client))

	return mux
}
