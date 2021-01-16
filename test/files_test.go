package test

import (
	"bufio"
	"os"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestReadDir(t *testing.T) {
	dirname := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log"

	f, err := os.Open(dirname)
	if err != nil {
		t.Error("Failed Open Directory")
	}

	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		t.Error("Failed Read Directory")
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	t.Log(list)
}

func TestReadFile(t *testing.T) {
	filename := "/Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log/http-2.log"
	minutes := 3

	f, err := os.Open(filename)
	if err != nil {
		t.Error("Failed Open File")
	}
	defer f.Close()

	result := []string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		r := strings.NewReplacer("[", "&;", "]", "&;")
		text := r.Replace(scanner.Text())
		text2 := strings.Split(text, "&;")

		tempTime := text2[1]
		layout := "02/Jan/2006:15:04:05 +0000"
		times, err := time.Parse(layout, tempTime)
		if err != nil {
			t.Error("Failed Parse Datetime")
		}

		now := time.Now()
		then := now.Add(time.Duration(-minutes) * time.Minute)
		if times.After(then) {
			result = append(result, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		t.Error("Failed Open File")
	}

	t.Log(result)
}
