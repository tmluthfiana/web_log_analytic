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

var (
	PathSeparator string = string(os.PathSeparator)
)

type LogAnalytic struct {
	Dirname  string
	Minute   int
	FileList []os.FileInfo
}

// this function use to read input from cli/command line
func Processes() error {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to Read Input")
		return err
	}

	stringTemp := strings.Fields(text)
	dirname := stringTemp[4]

	minTemp := strings.Replace(stringTemp[2], "m", "", -1)
	minute, err := strconv.Atoi(minTemp)
	if err != nil {
		fmt.Println("Failed to convert int")
		return err
	}

	var analytic = LogAnalytic{Dirname: dirname, Minute: minute}

	er := analytic.ProcessDir()
	if er != nil {
		fmt.Println("Failed to process dir")
		return er
	}

	return nil
}

//this function use to process dir(read and get list of log file in range n minutes)
func (analytic LogAnalytic) ProcessDir() error {
	now := time.Now()
	then := now.Add(time.Duration(-analytic.Minute) * time.Minute)

	fInfo, err := analytic.ReadDir()
	if err != nil {
		fmt.Println("Failed to Read Dir")
		return err
	}

	var files []os.FileInfo
	for _, file := range fInfo {
		if file.ModTime().After(then) {
			files = append(files, file)
		}
	}

	analytic.FileList = files
	er := analytic.ProcessFiles()
	if er != nil {
		fmt.Println("Failed to process file")
		return er
	}

	return nil
}

// this function used to process file (get n minutes information and save to temp store)
func (analytic LogAnalytic) ProcessFiles() error {

	if len(analytic.FileList) > 0 {
		fname := analytic.Dirname + PathSeparator + analytic.FileList[0].Name()
		err := analytic.CheckFirstFile(fname)
		if err != nil {
			fmt.Println("Failed to process file")
			return err
		}

		for _, file := range analytic.FileList {
			fname := analytic.Dirname + PathSeparator + file.Name()
			err := analytic.ReadFile(fname)
			if err != nil {
				fmt.Println("Failed to read file")
				return err
			}
		}
	}

	return nil
}

// this function used to oped and read the diresctory and save the information of the directory
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

// function for checking the first list of log file  n minutes
func (analytic LogAnalytic) CheckFirstFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to Open file")
		return err
	}
	defer f.Close()

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
			return err
		}

		now := time.Now()
		then := now.Add(time.Duration(-analytic.Minute) * time.Minute)
		if times.After(then) {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// this function use to oped and read the log file
func (analytic LogAnalytic) ReadFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to Open file")
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed read file")
		return err
	}

	return nil
}
