package routes

import (
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/handlers"
	"github.com/levionstudio/eddoswipe-backend/internal/middlewares"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (r *Router) walletRoutes(db *database.Database, jwtUtils *pkg.JWTUtils) {
	walletRepo := repositories.NewWalletRepository(db)
	walletHandler := handlers.NewWalletHandler(walletRepo)

	wrg := r.EchoRouter.Group("/wallets", middlewares.JWTMiddleware(jwtUtils))
	wrg.GET("/get/balance/admin/:admin_id", walletHandler.GetAdminWalletBalanceRequest)
	wrg.GET("/get/balance/md/:master_distributor_id", walletHandler.GetMasterDistributorWalletBalanceRequest)
	wrg.GET("/get/balance/distributor/:distributor_id", walletHandler.GetDistributorWalletBalanceRequest)
	wrg.GET("/get/balance/retailer/:retailer_id", walletHandler.GetRetailerWalletBalanceRequest)
	wrg.POST("/topup/admin", walletHandler.AdminWalletTopupRequest)
	wrg.GET("/get/ledger/entries/:id", walletHandler.GetLedgerTransactionsRequest)
}
