package objectValues

import "net/http"

type responseWithData struct {
	Title      string      `json:"title"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
}

func NewResponseWithData(StatusCode int, Title string, Message string, Data interface{}) responseWithData {
	return responseWithData{
		Title,
		Message,
		StatusCode,
		Data,
	}
}

type responseErrors struct {
	Title   string      `json:"title"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Errors  interface{} `json:"errors"`
}

func NewResponseErrors(Errors interface{}) responseErrors {
	return responseErrors{
		"Proceso no existoso",
		"Por favor corregir algunos errores",
		http.StatusBadRequest,
		Errors,
	}
}
