package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

func (db *Database) CreateCommisionQuery(ctx context.Context, req models.CreateCommisionModel) error {
	query := `
		INSERT INTO commisions (
			operator_id,
			operator_name,
			slab_start,
			slab_end,
			total_commision,
			admin_commision,
			master_distributor_commision,
			distributor_commision,
			retailer_commision
		) VALUES (
			@operator_id,
			@operator_name,
			@slab_start,
			@slab_end,
			@total_commision,
			@admin_commision,
			@master_distributor_commision,
			@distributor_commision,
			@retailer_commision 
		);
	`

	if _, err := db.pool.Exec(ctx, query, pgx.NamedArgs{
		"operator_id":                  req.OperatorID,
		"operator_name":                req.OperatorName,
		"slab_start":                   req.SlabStart,
		"slab_end":                     req.SlabEnd,
		"total_commision":              req.TotalCommision,
		"admin_commision":              req.AdminCommision,
		"master_distributor_commision": req.MasterDistributorCommision,
		"distributor_commision":        req.DistributorCommision,
		"retailer_commision":           req.RetailerCommision,
	}); err != nil {
		return fmt.Errorf("failed to create commision")
	}
	return nil
}

func (db *Database) GetAllCommisionsQuery(ctx context.Context) (*[]models.GetCommisionsModel, error) {
	query := `
		SELECT commision_id, operator_id, operator_name,
		slab_start, slab_end, total_commision, admin_commision,
		master_distributor_commision, distributor_commision,
		retailer_commision, created_at, updated_at
		FROM commisions;
	`
	res, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get commision data")
	}
	defer res.Close()

	var commisions []models.GetCommisionsModel
	for res.Next() {
		var commision models.GetCommisionsModel
		if err := res.Scan(
			&commision.CommisionID,
			&commision.OperatorID,
			&commision.OperatorName,
			&commision.SlabStart,
			&commision.SlabEnd,
			&commision.TotalCommision,
			&commision.AdminCommision,
			&commision.MasterDistributorCommision,
			&commision.DistributorCommision,
			&commision.RetailerCommision,
			&commision.CreatedAt,
			&commision.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get commision data")
		}

		commisions = append(commisions, commision)
	}
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get commision data")
	}
	return &commisions, nil
}

func (db *Database) UpdateCommisionQuery(ctx context.Context, req models.UpdateCommisionModel) error {
	query := `
		UPDATE commisions
		SET total_commision=@total_commision, admin_commision=@admin_commision,
		master_distributor_commision=@master_distributor_commision,
		distributor_commision=@distributor_commision
		WHERE commision_id=@commision_id;
	`
	if _, err := db.pool.Exec(ctx, query, pgx.NamedArgs{
		"total_commision":              req.TotalCommision,
		"admin_commision":              req.AdminCommision,
		"master_distributor_commision": req.MasterDistributorCommision,
		"distributor_commision":        req.DistributorCommision,
		"retailer_commision":           req.RetailerCommision,
		"commision_id":                 req.CommisionID,
	}); err != nil {
		return fmt.Errorf("failed to update commision")
	}
	return nil
}

func (db *Database) DeleteCommisionQuery(ctx context.Context, commisionID string) {
	
}
