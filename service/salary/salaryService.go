package salary

import (
	"github.com/namrahov/gross-to-net/model"
	"github.com/namrahov/gross-to-net/util"
	log "github.com/sirupsen/logrus"
)

type Service struct {
}

type IService interface {
	ConvertGrossToNet(dto *model.SalaryDto) float64
}

func (s *Service) ConvertGrossToNet(dto *model.SalaryDto) float64 {
	log.Info("ActionLog.GetEmployee.start")

	var taxAmount = dto.GrossSalary - dto.Discount

	var incomeTax = util.CalculateIncomeTax(taxAmount, dto.Sector)
	var dsmf = util.CalculateDsmf(dto.GrossSalary, dto.Sector)
	var unemploymentInsuranceFee = util.CalculateUnemploymentInsuranceFee(dto.GrossSalary)
	var membershipFee = util.CalculateMembershipFee(dto.GrossSalary, dto.ColleaguesFeePercent)
	var compulsoryHealthInsuranceFee = util.CalculateCompulsoryHealthInsuranceFee(dto.GrossSalary, dto.Sector)
	var deductedSalary = incomeTax + dsmf + unemploymentInsuranceFee + membershipFee + compulsoryHealthInsuranceFee

	log.Info("ActionLog.GetEmployee.success")

	return util.CalculateNetAmount(dto.GrossSalary, deductedSalary)
}
