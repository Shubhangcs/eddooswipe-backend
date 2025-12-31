package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/middlewares"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) authenticationRoutes(db *database.Database, jwtUtils *pkg.JWTUtils) {
	authRepo := repositories.NewAuthenticationRepository(db, jwtUtils)
	authHandler := handlers.NewAuthenticationHandler(authRepo)

	r.EchoRouter.POST("/auth/create/admin", authHandler.CreateAdminRequest)
	r.EchoRouter.POST("/auth/login/admin", authHandler.LoginAdminRequest)
	r.EchoRouter.POST("/auth/login/md", authHandler.LoginMasterDistributorRequest)
	r.EchoRouter.POST("/auth/login/distributor", authHandler.LoginDistributorRequest)
	r.EchoRouter.POST("/auth/login/retailer", authHandler.LoginRetailerRequest)

	arg := r.EchoRouter.Group("/auth", middlewares.JWTMiddleware(jwtUtils))
	arg.POST("/create/md", authHandler.CreateMasterDistributorRequest)
	arg.POST("/create/distributor", authHandler.CreateDistributorRequest)
	arg.POST("/create/retailer", authHandler.CreateRetailerRequest)

}
