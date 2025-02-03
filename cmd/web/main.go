package main

import (
	"log/slog"
	"net/http"
	"os"

	"web.herbalbones.com/internal/square"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	squareToken := os.Getenv("SQUARE_ACCESS_TOKEN")
	square := square.NewSquare(squareToken)

	serviceProvider := serviceProvider{
		squareService: square,
	}

	application := application{
		services: serviceProvider,
		logger:   logger,
	}

	http.ListenAndServe("0.0.0.0:4000", application.getRoutes())
}
