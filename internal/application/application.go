package application

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/m3owmurrr/calc/pkg/calc"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Run() error {
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
			log.Println(text, " calculation failed with error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}
