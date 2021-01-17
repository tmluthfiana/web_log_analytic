package test

import (
	"os"
	"testing"

	"github.com/tmluthfiana/web_log_analytic/api"
)

func TestProcessDir(t *testing.T) {
	dirname := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log"
	minute := 3
	var analytic = api.LogAnalytic{Dirname: dirname, Minute: minute}

	response, err := analytic.ProcessDir()
	t.Log(response)
	if err != nil {
		t.Error("Failed to Process Files")
	}
}

func TestProcessFiles(t *testing.T) {
	dirname := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log"
	minute := 10
	var analytic = api.LogAnalytic{Dirname: dirname, Minute: minute}

	filename := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log/http-2.log"
	info, err := os.Stat(filename)
	if err != nil {
		t.Error("file does not exist")
	}

	analytic.FileList = append(analytic.FileList, info)
	response, err := analytic.ProcessFiles()
	t.Log(response)
	if err != nil {
		t.Error("Failed to Process Files")
	}
}

func TestReadDir(t *testing.T) {
	dirname := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log"
	var analytic = api.LogAnalytic{Dirname: dirname}

	response, err := analytic.ReadDir()
	t.Log(response)
	if err != nil {
		t.Error("Failed to Read Directory")
	}
}

func TestReadFile(t *testing.T) {
	filename := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log/http-2.log"
	minute := 3
	var analytic = api.LogAnalytic{Minute: minute}

	response, err := analytic.ReadFile(filename)
	t.Log(response)
	if err != nil {
		t.Error("Failed to Read Files")
	}
}
