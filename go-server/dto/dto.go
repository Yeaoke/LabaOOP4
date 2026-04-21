package dto

type CompanyInfoDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Turnover    int64  `json:"turnover"`
	Type        string `json:"type"`
	HoldingName string `json:"holding_name,omitempty"`

	CoalVolume int64  `json:"coal_volume,omitempty"`
	MineCount  int64  `json:"mine_count,omitempty"`
	CoalAction string `json:"coal_action,omitempty"`

	OilVolume int64  `json:"oil_volume,omitempty"`
	WellCount int64  `json:"well_count,omitempty"`
	OilAction string `json:"oil_action,omitempty"`
}
