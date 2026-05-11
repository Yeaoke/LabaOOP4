package dto

import "github.com/google/uuid"

type IndustrialCompanies struct {
	ID             uuid.UUID `json:"id"`
	CompanyName    string    `json:"company_name,omitempty"  validate:"required,min=2"`
	AnnualTurnover int       `json:"annual_turnover,omitempty"`
	CompanyType    string    `json:"company_type,omitempty"  validate:"required,oneof=coal oil"`
	HoldingName    string    `json:"holding_name,omitempty"`

	Coal_volume int    `json:"coal_volume,omitempty"`
	MineCount   int    `json:"mineCount,omitempty"`
	Coal_action string `json:"coal_action,omitempty"`

	Oil_volume int    `json:"oil_volume,omitempty"`
	HoleCount  int    `json:"holeCount,omitempty"`
	Oil_action string `json:"oil_action,omitempty"`
}