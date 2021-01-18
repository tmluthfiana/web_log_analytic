package api

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	dirname := flag.String("dir", "foo", "to define directory of log files")
	minTemp := flag.String("t", "3m", "o define max duration of reading")
	flag.Parse()

	tempNum := strings.Replace(*minTemp, "m", "", -1)
	minute, err := strconv.Atoi(tempNum)
	if err != nil {
		fmt.Println("Failed to convert int")
		return err
	}

	var analytic = LogAnalytic{Dirname: *dirname, Minute: minute}
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

	fInfo, err := ioutil.ReadDir(analytic.Dirname)
	if err != nil {
		fmt.Println("Failed to Read Dir")
		return err
	}

	var files []os.FileInfo
	for _, file := range fInfo {
		if filepath.Ext(file.Name()) != ".log" {
			continue
		}

		if file.ModTime().After(then) {
			if !file.IsDir() {
				files = append(files, file)
			}
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
		for i, file := range analytic.FileList {
			fname := analytic.Dirname + PathSeparator + file.Name()
			if i == 0 {
				err := analytic.CheckFirstFile(fname)
				if err != nil {
					fmt.Println("Failed to process file")
					return err
				}
			} else {
				err := analytic.ReadFile(fname)
				if err != nil {
					fmt.Println("Failed to read file")
					return err
				}
			}
		}
	}

	return nil
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

		now := time.Now().UTC()
		then := now.Add(time.Duration(-analytic.Minute) * time.Minute)
		if times.After(then) {
			fmt.Println(scanner.Text())
		} else {
			continue
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
