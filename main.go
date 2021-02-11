package main

import (
	"github.com/golang/go/src/pkg/fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	usersCount, err := strconv.Atoi(args[1])
	if err != nil {
		panic("provide users count as a second argument")
	}

	logReader := InitLogReader(args[0], usersCount)
	timestamp, err := logReader.FindInterconnection()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Interconnection on %s", timestamp)
}