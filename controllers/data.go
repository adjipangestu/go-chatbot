package controllers

import (
	"net/http"
	"wa-chattbot/models"
	"wa-chattbot/responses"
)

func (server *Server) GetData(w http.ResponseWriter, r *http.Request)  {
	data := models.Indonesia{}

	datas, err := data.GetAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datas)
}