package routes

import (
	"fmt"
	"net/http"
)

const path = "users"

func handleUser(w http.ResponseWriter, r *http.Request) {

}

func SetupRoutes(apiBasePath string) {
	userHandler = http.HandlerFunc(handleUser)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, path), middleware.http.Middleware(userHandler))
}
