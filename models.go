package main

import (
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type tempDB struct {
	tasks  map[int]taskEntry
	mu     sync.RWMutex
	nextID int
	logger zerolog.Logger
}

type taskEntry struct {
	Id        int       `json:"id"`
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type parameters struct {
	Task string `json:"task"`
}
