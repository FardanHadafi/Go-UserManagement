package main

import (
	"Go-UserManagement/app"
	"Go-UserManagement/controller"
	"Go-UserManagement/helper"
	"Go-UserManagement/repository"
	"Go-UserManagement/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewCategoryController(userService)

	router := httprouter.New()
	router.POST("/api/register", userController.Register)
	router.POST("/api/login", userController.Login)
	
	server := http.Server{
		Addr: "localhost:3000",
		Handler: enableCORS(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}