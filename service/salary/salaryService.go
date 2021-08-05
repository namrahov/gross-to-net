package salary

import (
	"github.com/namrahov/gross-to-net/model"
	log "github.com/sirupsen/logrus"
)

type Service struct {
}

type IService interface {
	ConvertGrossToNet(dto *model.SalaryDto) float64
}

func (s *Service) ConvertGrossToNet(dto *model.SalaryDto) float64 {
	log.Info("ActionLog.GetEmployee.start")

	log.Info("ActionLog.GetEmployee.success")

	return dto.GrossSalary
}
