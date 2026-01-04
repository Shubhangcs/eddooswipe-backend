package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/middlewares"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) commisionRoutes(db *database.Database, jwtUtils *pkg.JWTUtils) {
	commisionRepo := repositories.NewCommisionRepository(db)
	commisionHandler := handlers.NewCommisionHandler(commisionRepo)

	crg := r.EchoRouter.Group("/commisions", middlewares.JWTMiddleware(jwtUtils))
	crg.POST("/create", commisionHandler.CreateCommisionRequest)
	crg.GET("/get/all", commisionHandler.GetAllCommisionsRequest)
	crg.PUT("/update", commisionHandler.UpdateCommisionRequest)
	crg.DELETE("/delete/:commision_id", commisionHandler.DeleteCommisionRequest)
}
