package test

import (
	"testing"

	"github.com/tmluthfiana/web_log_analytic/api"
)

func TestProcessFiles(t *testing.T) {
	dirname := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log"
	minutes := 3

	response, err := api.ProcessFiles(dirname, minutes)
	t.Log(response)
	if err != nil {
		t.Error("Failed to Process Files")
	}
}

func TestReadDir(t *testing.T) {
	dirname := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log"

	response, err := api.ReadDir(dirname)
	t.Log(response)
	if err != nil {
		t.Error("Failed to Read Directory")
	}
}

func TestReadFile(t *testing.T) {
	filename := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log/http-2.log"
	minutes := 3

	response, err := api.ReadFile(filename, minutes)
	t.Log(response)
	if err != nil {
		t.Error("Failed to Read Files")
	}
}
