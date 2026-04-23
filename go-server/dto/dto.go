package dto

type IndustrialCompaniesResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Turnover    int    `json:"turnover"`
	Type        string `json:"type"`
	HoldingName string `json:"holding_name,omitempty"`

	CoalVolume int    `json:"coal_volume,omitempty"`
	MineCount  int    `json:"mine_count,omitempty"`
	CoalAction string `json:"coal_action,omitempty"`

	OilVolume int    `json:"oil_volume,omitempty"`
	WellCount int    `json:"well_count,omitempty"`
	OilAction string `json:"oil_action,omitempty"`
}
