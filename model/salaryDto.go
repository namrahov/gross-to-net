package model

type SalaryDto struct {
	GrossSalary          float64 `json:"grossSalary"`
	Discount             float64 `json:"discount"`
	ColleaguesFeePercent int     `json:"colleaguesFeePercent"`
	Sector               Sector  `json:"sector"`
}

type Sector string

const (
	PRIVATE Sector = "PRIVATE"
	PUBLIC  Sector = "PUBLIC"
)
