package main

import (
	"os"
	"fmt"
	"time"
	"log"
	"strconv"
	"syscall"
)

const (
	STEP_SIZE = "100"
	MB_SIZE = 1024 * 1024
	PAGE_SIZE = 4096
)

func getenv(key, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}

func main() {
	step, err := strconv.Atoi(getenv("STEP_SIZE", STEP_SIZE))
	if err != nil {
		log.Fatal("STEP_SIZE should contain only numbers. Details: %s",err)
	}

	fmt.Printf("#### Allocate memory %d MiB every second ####\n", step)

	sum := 1
	for sum < 1000 {
		data, err := syscall.Mmap(-1, 0, step * MB_SIZE, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
		if err != nil {
			log.Printf("Failed to allocate memory %d MiB\n", sum *  step)
			log.Printf("Details: %s", err)
			break;
		}

		for i := 0; i < step * MB_SIZE; i += PAGE_SIZE {
			data[i] = 0
		}

		fmt.Printf("Allocated address = %p, Total Size = %d MB\n", &data[0], sum * step)
		time.Sleep(1 * time.Second)
		sum++

	}

	// wait forever
	for{}
}