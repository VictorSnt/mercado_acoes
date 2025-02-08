package DTO

type CreateEquitie struct {
	Name                  string
	CurrentPrince         float64
	PriceChangePercentage float64
}

type DisplayEquitie struct {
	ID                    uint    `json:"equitie_id"`
	Name                  string  `json:"name"`
	CurrentPrince         float64 `json:"current_price"`
	PriceChangePercentage float64 `json:"price_change_percentage"`
}

type UpdateEquitie struct {
	Name string `json:"name"`
}

type UpdateEquitiePrice struct {
	CurrentPrince float64
}

type UpdateUserBalance struct {
	Balance float64
}
