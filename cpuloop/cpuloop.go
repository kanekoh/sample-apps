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
	for{}
}

func main() {
	var wg sync.WaitGroup
	threadNum, err := strconv.Atoi(getenv("THREADS_NUM", THREADS_NUM))
	if err != nil {
		log.Fatal("THREADS_NUM should contain only numbers. Details: %s",err)
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