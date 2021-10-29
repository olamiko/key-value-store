package logging

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var COMMITLOG string = "commitlog.lg"

func setCommitLog(filename string) {
	COMMITLOG = filename
}

func writeToCommitLog(key string, value string) {
	f, err := os.OpenFile(COMMITLOG, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("commit log: " + err.Error())
	}
	defer f.Close()

	preparedInput := fmt.Sprintf("SET %s %s", key, value)
	if _, err := f.WriteString(preparedInput); err != nil {
		log.Fatal("commit log: " + err.Error())
	}
}

func rotateLog() {
	os.Remove(COMMITLOG)
}

func readCommitLog() *bufio.Scanner {
	f, err := os.Open(COMMITLOG)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return bufio.NewScanner(f)
	//	for scanner.Scan() {
	//		fmt.Println(scanner.Text())
	//	}

	//	if err := scanner.Err(); err != nil {
	//		log.Fatal(err)
	//	}
}
