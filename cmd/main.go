package main

import (
	"github.com/mdafaardiansyah/ptest-sltr-devops/internal/router"
	"log"
)

func main() {
	r := router.SetupRouter()

	if err := r.Run(":5000"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
