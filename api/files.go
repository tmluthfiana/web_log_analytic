package api

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ProcessDir() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to Read Dir")
	}

	stringTemp := strings.Fields(text)
	dirname := stringTemp[4]

	minTemp := strings.Replace(stringTemp[2], "m", "", -1)
	minutes, err := strconv.Atoi(minTemp)
	if err != nil {
		fmt.Println("Failed to convert int")
	}
	result := ProcessFiles(minutes, dirname)

	finalRes := strings.Join(result[:], "\n")
	return finalRes
}

func ProcessFiles(minutes int, dirname string) []string {
	now := time.Now()
	then := now.Add(time.Duration(-minutes) * time.Minute)

	fInfo, err := ReadDir(dirname)
	if err != nil {
		fmt.Println("Failed to Read Dir")
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
		res := ReadFile(filename, minutes)

		result = append(result, res...)
	}

	return result
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

func ReadFile(filename string, minutes int) []string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to Oper file")
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
		}

		now := time.Now()
		then := now.Add(time.Duration(-minutes) * time.Minute)
		if times.After(then) {
			result = append(result, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
