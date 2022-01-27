package models

import "time"

const BlocksInADay = 480

// Block represents a 5 minute block of time within a day.
type Block struct {
	Task  Task      `json:"task"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
