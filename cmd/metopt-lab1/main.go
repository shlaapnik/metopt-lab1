package main

import (
	"log"
	"metopt-lab1/internal/experiment"
)

func main() {
	if err := experiment.RunAllExperiments(); err != nil {
		log.Fatalf("Ошибка при выполнении экспериментов: %v", err)
	}
}
