package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/m3owmurrr/calc/internal/models"
)

func TestCalcHandler(t *testing.T) {
	testOK := []struct {
		name         string
		reqBody      models.ExpressionRequest
		wantStatus   int
		wantResponse models.ResultResponse
	}{
		{
			name:         "regular request 1",
			reqBody:      models.ExpressionRequest{Expression: "2+2"},
			wantStatus:   200,
			wantResponse: models.ResultResponse{Result: 4},
		},
		{
			name:         "regular request 2",
			reqBody:      models.ExpressionRequest{Expression: "2+2*2"},
			wantStatus:   200,
			wantResponse: models.ResultResponse{Result: 6},
		},
		{
			name:         "regular request 3",
			reqBody:      models.ExpressionRequest{Expression: "(2+2)*2"},
			wantStatus:   200,
			wantResponse: models.ResultResponse{Result: 8},
		},
		{
			name:         "regular request 4",
			reqBody:      models.ExpressionRequest{Expression: "8/2/2"},
			wantStatus:   200,
			wantResponse: models.ResultResponse{Result: 2},
		},
	}

	for _, tt := range testOK {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.reqBody)
			if err != nil {
				t.Fatalf("failed to marshal json: %v", err)
			}

			req := httptest.NewRequest("POST", "localhost:8080", bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			CalcHandler(w, req)
			resp := w.Result()

			if status := resp.StatusCode; status != tt.wantStatus {
				t.Errorf("CalcHandler for %v return %v status code, but want %v\n", tt.reqBody.Expression, status, tt.wantStatus)
			}

			var resResp models.ResultResponse
			if err := json.NewDecoder(resp.Body).Decode(&resResp); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			if resResp.Result != tt.wantResponse.Result {
				t.Errorf("CalcHandler for %v return result = %v, but want %v\n", tt.reqBody.Expression, resResp.Result, tt.wantResponse.Result)
			}
		})

	}

	testFail := []struct {
		name       string
		reqBody    models.ExpressionRequest
		wantStatus int
		wantError  models.ErrorResponse
	}{
		{
			name:       "unvalid expression 1",
			reqBody:    models.ExpressionRequest{Expression: "2^2"},
			wantStatus: 422,
			wantError:  models.ErrorResponse{Error: models.ErrNotValidExpression.Error()},
		},
		{
			name:       "unvalid expression 2",
			reqBody:    models.ExpressionRequest{Expression: "2+a"},
			wantStatus: 422,
			wantError:  models.ErrorResponse{Error: models.ErrNotValidExpression.Error()},
		},
		{
			name:       "unvalid expression 3",
			reqBody:    models.ExpressionRequest{Expression: "(2+2"},
			wantStatus: 422,
			wantError:  models.ErrorResponse{Error: models.ErrNotValidExpression.Error()},
		},
		{
			name:       "unvalid expression 4",
			reqBody:    models.ExpressionRequest{Expression: "2/0"},
			wantStatus: 422,
			wantError:  models.ErrorResponse{Error: models.ErrNotValidExpression.Error()},
		},
	}

	for _, tt := range testFail {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.reqBody)
			if err != nil {
				t.Fatalf("failed to marshal json: %v", err)
			}

			req := httptest.NewRequest("POST", "localhost:8080", bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			CalcHandler(w, req)
			resp := w.Result()

			if status := resp.StatusCode; status != tt.wantStatus {
				t.Errorf("CalcHandler for %v return %v status code, but want %v\n", tt.reqBody.Expression, status, tt.wantStatus)
			}

			var resResp models.ErrorResponse
			if err := json.NewDecoder(resp.Body).Decode(&resResp); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			if resResp.Error != tt.wantError.Error {
				t.Errorf("CalcHandler for %v return error \"%v\", but want \"%v\"\n", tt.reqBody.Expression, resResp.Error, tt.wantError.Error)
			}
		})

	}
}
