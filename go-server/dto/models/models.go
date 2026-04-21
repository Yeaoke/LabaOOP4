package models

import "github.com/google/uuid"

type IndustrialCompanies struct {
	ID             uuid.UUID `json:"id"`
	CompanyName    string    `json:"company_name"`
	AnnualTurnover int       `json:"turnover"`
	CompanyType    string    `json:"company_type"`
}

type CoalCompany struct {
	Coal_volume int    `json:"coal_volume"`
	MineCount   int    `json:"mineCount"`
	Coal_action string `json:"coal_action"`
}

type OilCompany struct {
	Oil_volume int    `json:"oil_volume"`
	HoleCount  int    `json:"holeCount"`
	Oil_action string `json:"oil_action"`
}

func NewCoalCompany() *CoalCompany {
	return &CoalCompany{}
}

func NewOilCompany() *OilCompany {
	return &OilCompany{
		//IndustrialCompanies: IndustrialCompanies{
		//	ID: uuid.New(),
		//},
	}
}
