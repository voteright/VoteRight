package models

import "time"

// IntegrityViolation represents a potential integrity violation
type IntegrityViolation struct {
	ID      int `storm:"id,increment"`
	Message string
	Time    time.Time
}
