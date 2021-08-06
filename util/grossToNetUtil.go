package util

import (
	"github.com/namrahov/gross-to-net/model"
)

func CalculateIncomeTax(taxAmount float64, sector model.Sector) float64 {
	var incomeTax float64

	switch sector {
	case model.PRIVATE:
		if taxAmount > 8000 {
			incomeTax = 0.14 * (taxAmount - 8000)
		} else {
			incomeTax = 0
		}
	case model.PUBLIC:
		if taxAmount <= 2500 {
			incomeTax = 0.14 * (taxAmount - 200)
		} else {
			incomeTax = 0.25*(taxAmount-2500) + 350
		}
	}

	return incomeTax
}

func CalculateDsmf(grossSalary float64, sector model.Sector) float64 {
	var dsmf float64

	switch sector {
	case model.PRIVATE:
		dsmf = 0.1*(grossSalary-200) + 6
	case model.PUBLIC:
		dsmf = 0.03 * grossSalary
	}

	return dsmf
}

func CalculateUnemploymentInsuranceFee(grossSalary float64) float64 {
	return 0.005 * grossSalary
}

func CalculateMembershipFee(grossSalary float64, colleaguesFeePercent int) float64 {
	return grossSalary * float64(colleaguesFeePercent/100)
}

func CalculateCompulsoryHealthInsuranceFee(grossSalary float64, sector model.Sector) float64 {
	var result float64

	switch sector {
	case model.PRIVATE:
		if grossSalary <= 8000 {
			result = 0.01 * grossSalary
		} else {
			result = 0.005*(grossSalary-8000) + 80
		}
	case model.PUBLIC:
		if grossSalary <= 8000 {
			result = 0.02 * grossSalary
		} else {
			result = 0.005*(grossSalary-8000) + 160
		}
	}

	return result
}

func CalculateNetAmount(grossSalary float64, deductedSalary float64) float64 {
	return grossSalary - deductedSalary
}
