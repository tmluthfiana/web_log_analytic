package api_test

import (
	"os"
	"testing"

	"github.com/tmluthfiana/web_log_analytic/api"
)

// test function that get list of lof file in n minutes
func TestProcessDir(t *testing.T) {
	dirname := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/http-log"
	minute := 3
	var analytic = api.LogAnalytic{Dirname: dirname, Minute: minute}

	err := analytic.ProcessDir()
	if err != nil {
		t.Error("Failed to Process Files")
	}
}

// test function that get n minutes data from log file
func TestProcessFiles(t *testing.T) {
	dirname := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/http-log"
	minute := 10
	var analytic = api.LogAnalytic{Dirname: dirname, Minute: minute}

	filename := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/http-log/http-2.log"
	info, err := os.Stat(filename)
	if err != nil {
		t.Error("file does not exist")
	}

	analytic.FileList = append(analytic.FileList, info)
	er := analytic.ProcessFiles()
	if er != nil {
		t.Error("Failed to Process Files")
	}
}

// test check first file
func TestCheckFirstFile(t *testing.T) {
	filename := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/http-log/http-2.log"
	minute := 3
	var analytic = api.LogAnalytic{Minute: minute}

	err := analytic.CheckFirstFile(filename)
	if err != nil {
		t.Error("Failed to Read Files")
	}
}

// this is function to test read a log file
func TestReadFile(t *testing.T) {
	filename := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/http-log/http-2.log"
	minute := 3
	var analytic = api.LogAnalytic{Minute: minute}

	err := analytic.ReadFile(filename)
	if err != nil {
		t.Error("Failed to Read Files")
	}
}
