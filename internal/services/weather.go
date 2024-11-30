package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Estrutura que guarda info do clima para a localidade
type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

/*
==========================================================
  - Função: GetWeatherByCity
  - Descrição : Função que busca dados do clima para a
  - localidade informada.
  - Parametros :
  - city - Localidade informada - tipo: string
  - Retorno: Informações do clima ou erro

==========================================================
*/
func GetWeatherByCity(city string) (*WeatherAPIResponse, error) {
	apiKey := "8d6494e11c514bb2882131719240511"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to get weather")
	}
	var weather WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}
	return &weather, nil
}