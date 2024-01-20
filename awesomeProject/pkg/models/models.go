package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord indicates that no matching record was found.
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrDuplicateRecord indicates a violation of a unique constraint.
	ErrDuplicateRecord = errors.New("models: duplicate record")
)

type News struct {
	ID       int
	Title    string
	Content  string
	Created  time.Time
	Expires  time.Time
	Category string
}
