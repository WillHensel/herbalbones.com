package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"web.herbalbones.com/internal/mailjet"
	"web.herbalbones.com/internal/square"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode

	squareToken := os.Getenv("SQUARE_ACCESS_TOKEN")
	square := square.NewSquare(squareToken)

	mailjetPub := os.Getenv("MAILJET_API_KEY")
	mailjetPriv := os.Getenv("MAILJET_SECRET_KEY")
	mailjet := mailjet.NewMailjet(mailjetPub, mailjetPriv)

	serviceProvider := serviceProvider{
		squareService:  square,
		mailjetService: mailjet,
	}

	application := application{
		services:       serviceProvider,
		logger:         logger,
		sessionManager: sessionManager,
	}

	http.ListenAndServe("0.0.0.0:4000", application.getRoutes())
}
