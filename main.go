package main

import (
	"fmt"

	"github.com/tmluthfiana/web_log_analytic/api"
)

func main() {
	err := api.Processes()
	if err != nil {
		fmt.Print("Failed to Process")
	}
}
