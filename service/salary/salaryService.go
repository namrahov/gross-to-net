package salary

import log "github.com/sirupsen/logrus"

type Service struct {
}

type IService interface {
	GetEmployee() float64
}

func (s *Service) GetEmployee() float64 {
	log.Info("ActionLog.GetEmployee.start")

	log.Info("ActionLog.GetEmployee.success")

	return 2.54
}
