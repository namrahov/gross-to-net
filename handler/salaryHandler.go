package handler

import (
	"github.com/go-chi/chi"
	"github.com/namrahov/gross-to-net/config"
	"github.com/namrahov/gross-to-net/service/salary"
	"net/http"
)

type employeeHandler struct {
	employeeService salary.IService
}

var employeeService = salary.Service{}

func EmployeeRequest(router *chi.Mux) *chi.Mux {
	h := &employeeHandler{employeeService: &employeeService}

	router.Get(config.ReportHandlerPath+"/employee/{id}", h.getEmployeeById)

	return router
}

func (h *employeeHandler) getEmployeeById(w http.ResponseWriter, r *http.Request) {

}
