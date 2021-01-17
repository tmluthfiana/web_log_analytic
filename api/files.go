package api

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type LogAnalytic struct {
	Dirname  string
	Minute   int
	FileList []os.FileInfo
}

func Processes() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to Read Input")
		return "", err
	}

	stringTemp := strings.Fields(text)
	dirname := stringTemp[4]

	minTemp := strings.Replace(stringTemp[2], "m", "", -1)
	minute, err := strconv.Atoi(minTemp)
	if err != nil {
		fmt.Println("Failed to convert int")
		return "", err
	}

	var analytic = LogAnalytic{Dirname: dirname, Minute: minute}

	result, err := analytic.ProcessDir()
	if err != nil {
		fmt.Println("Failed to process")
		return "", err
	}

	finalRes := strings.Join(result[:], "\n")
	return finalRes, nil
}

func (analytic LogAnalytic) ProcessDir() ([]string, error) {
	now := time.Now()
	then := now.Add(time.Duration(-analytic.Minute) * time.Minute)

	fInfo, err := analytic.ReadDir()
	if err != nil {
		fmt.Println("Failed to Read Dir")
		return nil, err
	}

	var files []os.FileInfo
	for _, file := range fInfo {
		if file.ModTime().After(then) {
			files = append(files, file)
		}
	}

	analytic.FileList = files
	result, err := analytic.ProcessFiles()
	if err != nil {
		fmt.Println("Failed to Process Files")
		return nil, err
	}

	return result, nil
}

func (analytic LogAnalytic) ProcessFiles() ([]string, error) {
	result := []string{}
	for _, file := range analytic.FileList {
		fmt.Println("name", file.Name())
		fname := analytic.Dirname + "/" + file.Name()
		res, err := analytic.ReadFile(fname)
		if err != nil {
			fmt.Println("Failed to Read File")
			return nil, err
		}

		result = append(result, res...)
	}
	return result, nil
}

func (analytic LogAnalytic) ReadDir() ([]os.FileInfo, error) {
	f, err := os.Open(analytic.Dirname)
	if err != nil {
		fmt.Println("Failed to Open Dir")
		return nil, err
	}

	list, err := f.Readdir(-1)
	if err != nil {
		fmt.Println("Failed to Read Dir")
		return nil, err
	}

	sort.Slice(list, func(i, j int) bool { return list[i].ModTime().Before(list[j].ModTime()) })
	f.Close()

	return list, nil
}

func (analytic LogAnalytic) ReadFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to Open file")
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
			fmt.Println("Failed to convert times")
			return nil, err
		}

		now := time.Now()
		then := now.Add(time.Duration(-analytic.Minute) * time.Minute)
		if times.After(then) {
			result = append(result, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
