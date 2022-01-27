package models

// Job represents a job to manage time for.
// (think a specific calendar)
type Job struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Hourly      bool    `json:"hourly"`
	HourlyRate  float64 `json:"hourly_rate"`
}
