package application

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/m3owmurrr/calc/internal/handlers"
	"github.com/m3owmurrr/calc/pkg/calc"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) RunLocal() error {
	for {
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}

		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("application was successfully")
			return nil
		}

		result, err := calc.Calc(text)
		if err != nil {
			log.Println(text, "calculation failed with error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

func (a *Application) RunServer() {
	m := http.NewServeMux()
	m.HandleFunc("POST /calculate", handlers.CalcHandler)
	m.HandleFunc("GET /health", handlers.HealthHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: m,
	}

	log.Println("server is running...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
