package models

type CreateCommisionModel struct {
	OperatorID                 string `json:"operator_id" validate:"required"`
	OperatorName               string `json:"operator_name" validate:"required"`
	SlabStart                  string `json:"slab_start" validate:"required"`
	SlabEnd                    string `json:"slab_end" validate:"required"`
	TotalCommision             string `json:"total_commision" validate:"required"`
	AdminCommision             string `json:"admin_commision" validate:"required"`
	MasterDistributorCommision string `json:"master_distributor_commision" validate:"required"`
	DistributorCommision       string `json:"distributor_commision" validate:"required"`
	RetailerCommision          string `json:"retailer_commision" validate:"required"`
}

type GetCommisionsModel struct {
	CommisionID                string `json:"commision_id"`
	OperatorID                 string `json:"operator_id"`
	OperatorName               string `json:"operator_name"`
	SlabStart                  string `json:"slab_start"`
	SlabEnd                    string `json:"slab_end"`
	TotalCommision             string `json:"total_commision"`
	AdminCommision             string `json:"admin_commision"`
	MasterDistributorCommision string `json:"master_distributor_commision"`
	DistributorCommision       string `json:"distributor_commision"`
	RetailerCommision          string `json:"retailer_commision"`
	CreatedAt                  string `json:"created_at"`
	UpdatedAt                  string `json:"updated_at"`
}

type UpdateCommisionModel struct {
	CommisionID                string `json:"commision_id" validate:"required"`
	TotalCommision             string `json:"total_commision" validate:"required"`
	AdminCommision             string `json:"admin_commision" validate:"required"`
	MasterDistributorCommision string `json:"master_distributor_commision" validate:"required"`
	DistributorCommision       string `json:"distributor_commision" validate:"required"`
	RetailerCommision          string `json:"retailer_commision" validate:"required"`
}
