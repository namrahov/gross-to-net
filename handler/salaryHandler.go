package handler

import (
	"github.com/go-chi/chi"
	"github.com/namrahov/gross-to-net/config"
)

func convertGrossToNet(router *chi.Mux) *chi.Mux {
	h := &employeeHandler{permissionService: &permissionService, employeeService: &employeeService}

	router.Get(config.ReportHandlerPath+"/employee/{id}", h.getEmployeeById)

	return router
}
