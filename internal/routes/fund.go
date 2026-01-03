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
	frg.POST("/create/fund/request/admin/md", fundHandler.CreateMasterDistributorFundRequestAdminRequest)
	frg.POST("/create/fund/request/admin/distributor", fundHandler.CreateDistributorFundRequestAdminRequest)
	frg.POST("/create/fund/request/admin/retailer", fundHandler.CreateRetailerFundRequestAdminRequest)
	frg.POST("/accept/fund/request", fundHandler.AcceptFundRequest)
	frg.POST("/reject/fund/request", fundHandler.RejectFundRequest)
	frg.GET("/get/fund/request/to/admin/:id", fundHandler.GetRequestToFundRequest)
	frg.GET("/get/fund/request/to/md/:id", fundHandler.GetRequestToFundRequest)
	frg.GET("/get/fund/request/to/distributor/:id", fundHandler.GetRequestToFundRequest)
	frg.GET("/get/fund/request/from/md/:id", fundHandler.GetRequesterFundRequest)
	frg.GET("/get/fund/request/from/distributor/:id", fundHandler.GetRequesterFundRequest)
	frg.GET("/get/fund/request/from/retailer/:id", fundHandler.GetRequesterFundRequest)
}
