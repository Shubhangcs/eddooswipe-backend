package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) dmtRouter(db *database.Database, jwtUtils *pkg.JWTUtils) {
	dmtRepo := repositories.NewDMTRepository(db,jwtUtils)
	dmtHandler := handlers.NewDMTHandler(dmtRepo)

	r.EchoRouter.GET("/dmt/register", dmtHandler.RegisterMerchantRequest)
}
