package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mr-isik/go-fred-scraper/internal/client"
	"github.com/mr-isik/go-fred-scraper/internal/series"
	"github.com/mr-isik/go-fred-scraper/internal/writer"
	"github.com/mr-isik/go-fred-scraper/pkg/config"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Configuration
	cfg := config.New()

	// Flags
	seriesID := flag.String("series", "", "FRED series ID (required)")
	startDate := flag.String("start", "", "Start date in YYYY-MM-DD format (required)")
	endDate := flag.String("end", "", "End date in YYYY-MM-DD format (required)")
	outputFile := flag.String("output", "output.csv", "Output CSV file path")
	flag.Parse()

	if *seriesID == "" || *startDate == "" || *endDate == "" {
		flag.Usage()
		os.Exit(1)
	}

	if cfg.APIKey == "" {
		log.Fatal("FRED_API_KEY environment variable not set")
	}

	// Create FRED client
	fredClient := client.New(cfg)

	// Create series service
	seriesService := series.NewService(fredClient)

	// Get series observations
	obs, err := seriesService.GetSeriesObservations(*seriesID, *startDate, *endDate)
	if err != nil {
		log.Fatalf("Error getting series observations: %v", err)
	}

	// Create CSV writer
	csvWriter := writer.NewCSVWriter(*outputFile)

	// Write to CSV
	if err := csvWriter.Write(obs); err != nil {
		log.Fatalf("Error writing to CSV: %v", err)
	}

	fmt.Printf("Successfully wrote data to %s\n", *outputFile)
}
