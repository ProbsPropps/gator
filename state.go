package main

import (
	"github.com/ProbsPropps/gator/internal/config"
	"github.com/ProbsPropps/gator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}
