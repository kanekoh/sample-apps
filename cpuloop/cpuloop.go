package main

import (
	"os"
	"strconv"
	"log"
	"fmt"
	"sync"
)

const (
	THREADS_NUM = "4"
	LOOP_VAR = "10000000000"
)

var (
	loopVar = 0
)

func getenv(key, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}

func forever(wg *sync.WaitGroup) {
	defer wg.Done()
	// wait forever
	for i := 0; i < loopVar; i++ {}
}

func main() {
	var wg sync.WaitGroup
	threadNum, err := strconv.Atoi(getenv("THREADS_NUM", THREADS_NUM))
	if err != nil {
		log.Fatal("THREADS_NUM should contain only numbers. Details: %s",err)
	}
	loopVar, err = strconv.Atoi(getenv("LOOP_VAR", LOOP_VAR))
	if err != nil {
		log.Fatal("LOOP_VAR should contain only numbers. Details: %s",err)
	}

	fmt.Println("#### Start loop ####")

	for i := 0; i < threadNum; i++ {
		fmt.Printf("Start thread. id = %d\n", i)
		wg.Add(1)
		go forever(&wg)
	}

	// wait to finish all threads
	wg.Wait()
}