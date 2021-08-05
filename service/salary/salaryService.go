package salary

import (
	log "github.com/sirupsen/logrus"
	"math/big"
)

type Service struct {
}

type IService interface {
	ConvertGrossToNet() big.Int
}

func (s *Service) ConvertGrossToNet() float64 {
	log.Info("ActionLog.GetEmployee.start")

	log.Info("ActionLog.GetEmployee.success")

	return 2.4
}
