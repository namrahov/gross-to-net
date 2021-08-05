package salary

import (
	log "github.com/sirupsen/logrus"
)

type Service struct {
}

type IService interface {
	ConvertGrossToNet(id int) (*model.EmployeeDto, *model.ErrorResponse)
}

func (s *Service) ConvertGrossToNet(id int) (*model.EmployeeDto, *model.ErrorResponse) {
	log.Info("ActionLog.GetEmployee.start")

	employeeList := []model.EmployeeDto{
		{1, `Nurlan`},
		{2, `Hesen`},
	}

	for _, v := range employeeList {
		if v.Id == id {
			return &v, nil
		}
	}

	log.Info("ActionLog.GetEmployee.success")

	return nil, nil
}
