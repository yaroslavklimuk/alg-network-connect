package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type (
	LogReader struct {
		connTree	ConnectionTree
		file		string
	}
)

func InitLogReader(file string, usersCount int) *LogReader {
	return &LogReader{
		connTree: InitConnectionTree(usersCount),
		file:     file,
	}
}

func (rdr *LogReader) FindInterconnection() (string, error) {
	fileHndl, err := os.Open(rdr.file)
	if err != nil {
		return "", errors.New("can't open a file")
	}
	defer fileHndl.Close()

	scanner := bufio.NewScanner(fileHndl)
	for scanner.Scan() {
		lineStr := scanner.Text()
		lineParts := strings.Split(lineStr, ";")
		if len(lineParts) != 3 {
			return "", errors.New("bad log format")
		}
		left, err := strconv.Atoi(lineParts[1])
		if err != nil {
			return "", err
		}
		right, err := strconv.Atoi(lineParts[2])
		if err != nil {
			return "", err
		}
		_, err = rdr.connTree.Union(left, right)
		if err != nil {
			return "", err
		}
		if rdr.connTree.IsInterconnected() {
			return lineParts[0], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}
