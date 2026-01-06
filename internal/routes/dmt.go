package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) dmtRouter(db *database.Database, jwtUtils *pkg.JWTUtils, paysprintURL string) {
	dmtRepo := repositories.NewDMTRepository(db, jwtUtils, paysprintURL)
	dmtHandler := handlers.NewDMTHandler(dmtRepo)

	r.EchoRouter.GET("/dmt/check/merchant/:retailer_id", dmtHandler.CheckMerchantRegistrationRequest)
	r.EchoRouter.POST("/dmt/register/merchant", dmtHandler.RegisterMerchantRequest)
}
