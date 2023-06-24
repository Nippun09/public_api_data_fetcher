package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"workspace/config"
	"workspace/constants"
	"workspace/core/fetcher"
)

func main() {
	// Load configuration from file
	config, err := config.LoadConfig(constants.CONFIG_FILE_PATH)
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	// Create ticker to periodically fetch data
	ticker := time.NewTicker(config.Frequency)
	defer ticker.Stop()

	// Create a channel to receive OS signals, part of shutdown hook for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Create a channel to indicate when to stop fetching data
	stopFetching := make(chan struct{})

	// Start a goroutine to handle OS signals. this is actual shutdown hook
	go func() {
		<-signals
		fmt.Println("Received shutdown signal. Stopping data fetching...")
		// Handle all cleanup here if any
		close(stopFetching)
	}()

	// Fetch data on initial run
	fetcher.FetchData(config.OutputFile)

	// Loop to periodically fetch data until stop signal is received
	for {
		select {
		case <-ticker.C:
			fetcher.FetchData(config.OutputFile)
		case <-stopFetching:
			fmt.Println("Stopping data fetching...")
			return
		}
	}
}
