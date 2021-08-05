package handler

import (
	"github.com/go-chi/chi"
	"github.com/namrahov/gross-to-net/config"
	"github.com/namrahov/gross-to-net/service/salary"
	"net/http"
)

type salaryHandler struct {
	salaryService salary.IService
}

var employeeService = salary.Service{}

func ConvertGrossToNet(router *chi.Mux) *chi.Mux {
	h := &salaryHandler{salaryService: &employeeService}

	router.Get(config.ReportHandlerPath+"/employee/{id}", h.ConvertGrossToNet)

	return router
}

func (h *salaryHandler) ConvertGrossToNet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(employeeDto)
}
