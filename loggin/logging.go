package loggin

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var increasingStamp int = 1

var COMMITLOG string = "commitlog.lg"

func SetCommitLog(filename string) {
	COMMITLOG = filename
}

func WriteCommitLog(key string, value string) {
	f, err := os.OpenFile(COMMITLOG, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("commit log: " + err.Error())
	}
	defer f.Close()

	preparedInput := fmt.Sprintf("%d SET %s %s\n", increasingStamp, key, value)
	if _, err := f.WriteString(preparedInput); err != nil {
		log.Fatal("commit log: " + err.Error())
	}
	increasingStamp = increasingStamp + 1
}

func RotateLog() {
	os.Remove(COMMITLOG)
}

func ReadCommitLog() []string {

	commitSlice := make([]string, 0)

	f, err := os.Open(COMMITLOG)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		commitSlice = append(commitSlice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return commitSlice
	//	for scanner.Scan() {
	//		fmt.Println(scanner.Text())
	//	}

	//	if err := scanner.Err(); err != nil {
	//		log.Fatal(err)
	//	}
}
