package app

import (
	"banking-app/dto"
	"banking-app/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ch *AccountHandlers) save(w http.ResponseWriter, r *http.Request) {
	var dto dto.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Fprint(w, "Failed to parse the body")
	}

	responseDto, err2 := ch.service.Save(dto)
	if err2 != nil {
		writeResponse(w, err2.Code, err2)
	} else {
		writeResponse(w, http.StatusOK, responseDto)
	}

}
