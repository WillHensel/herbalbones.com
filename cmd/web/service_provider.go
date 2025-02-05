package main

import (
	"web.herbalbones.com/internal/mailjet"
	"web.herbalbones.com/internal/square"
)

type serviceProvider struct {
	squareService  square.Square
	mailjetService mailjet.Mailjet
}
