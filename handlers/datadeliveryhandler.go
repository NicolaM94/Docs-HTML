package handlers

import (
	"net/http"
)

func DataDeliveryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reached data delivery"))
}
