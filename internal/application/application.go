package application

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/m3owmurrr/calc/internal/config"
	"github.com/m3owmurrr/calc/internal/handlers"
	"github.com/m3owmurrr/calc/pkg/calc"
)

type Application struct {
	Config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{
		Config: config,
	}
}

func (a *Application) Run() {
	if a.Config.RunType == "Local" {
		a.RunLocal()
	} else if a.Config.RunType == "Server" {
		a.RunServer()
	}
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

	addr := fmt.Sprintf("%v:%v", a.Config.Host, a.Config.Port)
	server := http.Server{
		Addr:    addr,
		Handler: m,
	}

	log.Printf("server is running on %v...", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
