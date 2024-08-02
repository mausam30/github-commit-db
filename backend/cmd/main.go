package main

import (
	"github-commit-viewer/pkg/api"
	"github-commit-viewer/pkg/utils"
	"log"
)

func main() {
	utils.LoadEnv()
	router := api.SetupRouter()
	log.Fatal(router.Run(":8080"))
}
