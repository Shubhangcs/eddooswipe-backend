package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/middlewares"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) fundRoutes(db *database.Database, jwtUtils *pkg.JWTUtils) {
	fundRepo := repositories.NewFundRepository(db)
	fundHandler := handlers.NewFundHandler(fundRepo)

	frg := r.EchoRouter.Group("/funds", middlewares.JWTMiddleware(jwtUtils))
	frg.POST("/crearte/fund/request", fundHandler.CreateFundRequest)
	frg.POST("/accept/fund/request", fundHandler.AcceptFundRequest)
	frg.POST("/reject/fund/request", fundHandler.RejectFundRequest)
}
