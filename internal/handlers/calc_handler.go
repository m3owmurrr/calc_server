package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/m3owmurrr/calc/internal/models"
	"github.com/m3owmurrr/calc/pkg/calc"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // закрываем реквест

	var expr models.ExpressionRequest
	if err := json.NewDecoder(r.Body).Decode(&expr); err != nil {
		errResp := models.ErrorResponse{
			Error: models.ErrNotValidJson.Error(),
		}

		// отправляем не 500, т.к. проблема пришла со стороны пользователя
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	result, err := calc.Calc(expr.Expression)
	if err != nil {
		errResp := models.ErrorResponse{
			Error: models.ErrNotValidExpression.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	resResp := models.ResultResponse{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resResp)
}
