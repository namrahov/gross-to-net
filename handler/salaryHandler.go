package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/namrahov/gross-to-net/config"
	"github.com/namrahov/gross-to-net/errhandler"
	"github.com/namrahov/gross-to-net/model"
	"github.com/namrahov/gross-to-net/service/salary"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type salaryHandler struct {
	salaryService salary.IService
}

var salaryService = salary.Service{}

func ConvertGrossToNet(router *chi.Mux) *chi.Mux {
	h := &salaryHandler{salaryService: &salaryService}

	router.Get(config.ReportHandlerPath+"/employee/{id}", h.ConvertGrossToNet)

	return router
}

func (h *salaryHandler) ConvertGrossToNet(w http.ResponseWriter, r *http.Request) {

	dto, err := getRatingFromRequest(r)
	if err != nil {
		return
	}

	result := new(float64)

	*result = h.salaryService.ConvertGrossToNet(dto)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getRatingFromRequest(r *http.Request) (*model.SalaryDto, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("failed to parse the model", b)
		return nil, errhandler.NewBadRequestError(err.Error(), nil)
	}

	dto := new(model.SalaryDto)
	err = json.Unmarshal(b, dto)
	if err != nil {
		log.Error("failed to parse the model", b)
		return nil, errhandler.NewBadRequestError(err.Error(), nil)
	}

	return dto, nil
}
