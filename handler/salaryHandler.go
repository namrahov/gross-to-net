package handler

import (
	"github.com/go-chi/chi"
	"github.com/namrahov/gross-to-net/config"
	"github.com/namrahov/gross-to-net/salary"
)

type salaryHandler struct {
	salaryService salary.IService
}

var salaryService = salary.Service{}

func convertGrossToNet(router *chi.Mux) *chi.Mux {
	h := &salaryHandler{salaryService: &salaryService}

	router.Get(config.ReportHandlerPath+"/employee/{id}", h.getEmployeeById)

	return router
}
