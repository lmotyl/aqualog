package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gihub.com/lmotyl/aqualog/middleware"
)

const path = "users"

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		response, err := json.Marshal()
	}
}

func SetupRoutes(apiBasePath string) {
	userHandler := http.HandlerFunc(handleUser)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, path), middleware.Middleware(userHandler))
}
