package main

import (
	"fmt"

	"github.com/tmluthfiana/web_log_analytic/api"
)

func main() {
	result, err := api.Processes()
	if err != nil {
		fmt.Print("Failed to Process")
	} else {
		fmt.Print(result)
	}
}
