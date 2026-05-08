package routes

import (
	"mvc-orm/internal/controllers"
	"mvc-orm/internal/middlewares"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	publicRegisterHandler := middlewares.LoggerMiddleware(controllers.RegisterUser)
	mux.HandleFunc("POST /api/v1/users/register", publicRegisterHandler)
}