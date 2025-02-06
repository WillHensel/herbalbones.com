package main

import "log/slog"

type application struct {
	services serviceProvider
	logger   *slog.Logger
}
