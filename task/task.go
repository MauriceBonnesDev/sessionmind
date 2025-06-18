package task

import (
	"time"
)

type Task struct {
	ID 				int
	Name 			string
	Description 	string
	CreatedAt 		time.Time
	StartedAt 		*time.Time
	EndedAt 		*time.Time
	Duration 		time.Duration
	Done 			bool
	Paused			bool
	Active 			bool
}
