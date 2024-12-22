package handlers

import (
	"fmt"
	"net/http"
)

// HealthHandler используется для проверки состояния сервера
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "It's alive!\n")
}
