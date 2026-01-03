package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/middlewares"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) bankRouter(db *database.Database, jwtUtils *pkg.JWTUtils) {
	bankRepo := repositories.NewBankRepository(db)
	bankHandler := handlers.NewBankHandler(bankRepo)

	brg := r.EchoRouter.Group("/banks", middlewares.JWTMiddleware(jwtUtils))
	brg.POST("/create/bank/admin", bankHandler.CreateAdminBankRequest)
	brg.POST("/create/bank/retailer", bankHandler.CreateRetailerBankRequest)
	brg.GET("/get/all/admin", bankHandler.GetAllAdminBanksRequest)
	brg.GET("/get/all/retailer", bankHandler.GetAllRetailerBanksRequest)
	brg.GET("/get/retailer/:retailer_id", bankHandler.GetRetailerBanksByRetailerIDRequest)
	brg.GET("/get/admin/:admin", bankHandler.GetAdminBanksByAdminIDRequest)
}
