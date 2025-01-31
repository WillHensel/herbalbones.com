package main

import (
	"net/http"
	"os"

	"web.herbalbones.com/internal/square"
)

func main() {

	squareToken := os.Getenv("SQUARE_ACCESS_TOKEN")
	square := square.NewSquare(squareToken)

	serviceProvider := serviceProvider{
		squareService: square,
	}

	application := application{
		services: serviceProvider,
	}

	http.ListenAndServe("localhost:4000", application.getRoutes())
}
