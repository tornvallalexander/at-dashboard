package main

import (
	"fmt"
	"log"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckFatalError(err error) {
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %s\n", err)
	}
}
