package database

import (
	"context"
	"fmt"

	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

func (db *Database) CreateOperatorQuery(ctx context.Context, operatorName string) error {
	query := `
		INSERT INTO operators (
			operator_name
		) VALUES (
			@operator_name 
		);
	`
	if _, err := db.pool.Exec(ctx, query, operatorName); err != nil {
		return fmt.Errorf("failed to create operator")
	}
	return nil
}

func (db *Database) GetAllOperators(ctx context.Context) (*[]models.GetOperatorsModel, error) {
	query := `
		SELECT operator_id, operator_name 
		FROM operators;
	`
	res, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get operators data")
	}
	defer res.Close()

	var operators []models.GetOperatorsModel
	for res.Next() {
		var operator models.GetOperatorsModel
		if err := res.Scan(&operator.OperatorID, &operator.OperatorName); err != nil {
			return nil, fmt.Errorf("failed to get operators data")
		}
		operators = append(operators, operator)
	}

	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get operators data")
	}
	return &operators, nil
}
