package main

import (
	"github.com/alexandrecsmoreira/gopportunities/config"
	"github.com/alexandrecsmoreira/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Debugf("Error xpto %v", err)
		return
	}

	router.Initialize()
}
