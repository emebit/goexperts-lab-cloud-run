package handler

import (
	"Labs/goexperts-lab-cloud-run/internal/models"
	"Labs/goexperts-lab-cloud-run/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

/*
==========================================================
  - Função: isValidCEP
  - Descrição : Função que verifica se o CEP é válido
  - Parametros :
  - zipCode - CEP - tipo: string
  - Retorno: Booleano

==========================================================
*/
func isValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

/*
==========================================================
  - Função: CepHandler
  - Descrição : Função que valida e busca dados do clima
  - CEP informado.
  - Parametros :
  - w - Resposta do HTTP - tipo: http.ResponseWriter
  - r - Ponteiro para a requisição do HTTP tipo: http.Request
  - Retorno: Informações do clima no Response HTTP

==========================================================
*/
func CepHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cep := strings.TrimPrefix(r.URL.Path, "/cep/")
	if !isValidCEP(cep) {
		http.Error(w, `{"message": "invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}
	fmt.Println(cep)
	location, err := services.GetLocationByCEP(cep)
	if err != nil {
		http.Error(w, `{"message": "can not found zipcode"}`, http.StatusNotFound)
		return
	}
	weather, err := services.GetWeatherByCity(location.City)
	if err != nil {
		http.Error(w, `{"message": "could not retrieve weather"}`, http.StatusInternalServerError)
		return
	}
	tempF := weather.Current.TempC*1.8 + 32
	tempK := weather.Current.TempC + 273.15
	response := models.TemperatureResponse{
		TempC: weather.Current.TempC,
		TempF: tempF,
		TempK: tempK,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
