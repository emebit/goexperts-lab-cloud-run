package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Estrutura que guarda a localidade do CEP
type ViaCEPResponse struct {
	City string `json:"localidade"`
}

/*
==========================================================
  - Função: GetlocationByCEP
  - Descrição : Função que busca dados da localidade para o
  - CEP informado.
  - Parametros :
  - cep - CEP informado tipo: string
  - Retorno: Informações da localidade ou erro

==========================================================
*/
func GetLocationByCEP(cep string) (*ViaCEPResponse, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("Error fetching data from ViaCEP:", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("ViaCEP API returned non-OK status:", resp.StatusCode)
		return nil, fmt.Errorf("unable to get location")
	}
	var location ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		log.Println("Error decoding ViaCEP response:", err)
		return nil, err
	}
	if location.City == "" {
		log.Printf("CEP %s not found in ViaCEP response.\n", cep)
		return nil, fmt.Errorf("location not found")
	}
	log.Printf("ViaCEP response for CEP %s: %+v\n", cep, location)
	return &location, nil
}
