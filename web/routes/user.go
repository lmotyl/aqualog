package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gihub.com/lmotyl/aqualog/middleware"
	"gihub.com/lmotyl/aqualog/user"
)

const path = "users"

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("get method")
		user := user.CreateUser(1, "Coala Bear", 145678)
		response, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(response)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SetupRoutes(apiBasePath string) {
	userHandler := http.HandlerFunc(handleUser)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, path), middleware.Middleware(userHandler))
}
