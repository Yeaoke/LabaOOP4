package models

import "github.com/google/uuid"

type IndustrialCompanies struct {
	ID             uuid.UUID `json:"id"`
	CompanyName    string    `json:"company_name,omitempty" validate:"required,min=2"`
	AnnualTurnover int       `json:"annual_turnover,omitempty" validate:"min=0"`
	CompanyType    string    `json:"company_type,omitempty" validate:"required,oneof=coal oil"`
	HoldingName    string    `json:"holding_name,omitempty"`

	Coal_volume int    `json:"coal_volume,omitempty" validate:"min=0"`
	MineCount   int    `json:"mineCount,omitempty" validate:"min=0"`
	Coal_action string `json:"coal_action,omitempty"`

	Oil_volume int    `json:"oil_volume,omitempty" validate:"min=0"`
	HoleCount  int    `json:"holeCount,omitempty" validate:"min=0"`
	Oil_action string `json:"oil_action,omitempty"`
}