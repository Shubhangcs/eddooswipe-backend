package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/middlewares"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) OperatorRoutes(db *database.Database, jwtUtils *pkg.JWTUtils) {
	operatorRepo := repositories.NewOperatorRepository(db)
	operatorHandler := handlers.NewOperatorHandler(operatorRepo)

	org := r.EchoRouter.Group("/operators", middlewares.JWTMiddleware(jwtUtils))
	org.POST("/create/operator", operatorHandler.CreateOperatorRequest)
	org.GET("/get/all", operatorHandler.GetAllOperatorsRequest)
}
