package models

import "github.com/google/uuid"

type IndustrialCompanies struct {
	ID             uuid.UUID `json:"id" validate:"required,uuid"`
	CompanyName    string    `json:"company_name,omitempty" validate:"required,min=2"`
	AnnualTurnover int       `json:"turnover,omitempty" validate:"min=0"`
	CompanyType    string    `json:"company_type,omitempty" validate:"required,oneof=coal oil"`
	HoldingName    string    `json:"holding_name,omitempty" validate:"required,min=2"`

	Coal_volume int    `json:"coal_volume,omitempty" validate:"required,min=0"`
	MineCount   int    `json:"mineCount,omitempty" validate:"required,min=1"`
	Coal_action string `json:"coal_action,omitempty"`

	Oil_volume int    `json:"oil_volume,omitempty" validate:"required,min=0"`
	HoleCount  int    `json:"holeCount,omitempty" validate:"required,min=1"`
	Oil_action string `json:"oil_action,omitempty"`
}

func NewCoalCompany() *IndustrialCompanies {
	return &IndustrialCompanies{
		//	ID: uuid.New(),
	}
}
