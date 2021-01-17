package api

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ProcessDir() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	stringTemp := strings.Fields(text)
	dirname := stringTemp[4]

	minTemp := strings.Replace(stringTemp[2], "m", "", -1)
	minutes, err := strconv.Atoi(minTemp)
	if err != nil {
		return "", err
	}

	result, err := ProcessFiles(dirname, minutes)
	if err != nil {
		return "", err
	}

	finalRes := strings.Join(result[:], "\n")
	return finalRes, nil
}

func ProcessFiles(dirname string, minutes int) ([]string, error) {
	now := time.Now()
	then := now.Add(time.Duration(-minutes) * time.Minute)

	fInfo, err := ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	var files []os.FileInfo
	for _, file := range fInfo {
		if file.ModTime().After(then) {
			files = append(files, file)
		}
	}

	result := []string{}
	for _, file := range files {
		filename := dirname + "/" + file.Name()
		res, err := ReadFile(filename, minutes)
		if err != nil {
			return nil, err
		}

		result = append(result, res...)
	}

	return result, nil
}

func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

func ReadFile(filename string, minutes int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		now := time.Now()
		then := now.Add(time.Duration(-minutes) * time.Minute)
		if times.After(then) {
			result = append(result, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
