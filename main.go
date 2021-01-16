package main

import (
	"fmt"

	"github.com/tmluthfiana/web_log_analytic/api"
)

func main() {
	result := api.ProcessDir()
	fmt.Print(result)
}
