package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/middlewares"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) usersRoutes(db *database.Database, jwtUtils *pkg.JWTUtils) {
	userRepo := repositories.NewUsersRepository(db, jwtUtils)
	userHandler := handlers.NewUsersHandler(userRepo)
	urg := r.EchoRouter.Group("/users", middlewares.JWTMiddleware(jwtUtils))
	urg.GET("/get/all/admins", userHandler.GetAllAdminsRequest)
	urg.GET("/get/all/mds", userHandler.GetAllMasterDistributorsRequest)
	urg.GET("/get/all/distributors", userHandler.GetAllDistributorsRequest)
	urg.GET("/get/all/retailers", userHandler.GetAllRetailersRequest)
	urg.GET("/get/mds/:admin_id", userHandler.GetMasterDistributorsByAdminIDRequest)
	urg.GET("/get/distributors/:master_distributor_id", userHandler.GetDistributorsByMasterDistributorIDRequest)
	urg.GET("/get/retailers/:distributor_id", userHandler.GetRetailersByDistributorIDRequest)
	urg.PATCH("/block/md/:master_distributor_id", userHandler.BlockMasterDistributorRequest)
	urg.PATCH("/block/distributor/:distributor_id", userHandler.BlockDistributorRequest)
	urg.PATCH("/block/retailer/:retailer_id", userHandler.BlockRetailerRequest)
	urg.PATCH("/unblock/md/:master_distributor_id", userHandler.UnBlockMasterDistributorRequest)
	urg.PATCH("/unblock/distributor/:distributor_id", userHandler.UnBlockDistributorRequest)
	urg.PATCH("/unblock/retailer/:retailer_id", userHandler.UnBlockRetailerRequest)
}
