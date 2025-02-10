package main

import (
	"log/slog"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	services       serviceProvider
	sessionManager *scs.SessionManager
	logger         *slog.Logger
}
