package Domain

import 
"time"

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Date time.Time `json:"date"`
	Category string `json:"category"`
	Status string `json:"status"`
}