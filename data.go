package goClassData

import (
	"time"
)

type Employees map[string]Employee
type Employee struct {
	FirstName       string    `json:"FirstName"`
	LastName        string    `json:"LastName"`
	StartDate       time.Time `json:"StartDate"`
	TerminationDate time.Time `json:"TerminationDate"`
	Active          bool      `json:"Active"`
}

func (e *Employee) IsActive() (active bool) {
	active = false

	if e.Active {
		if time.Now().Before(e.TerminationDate) {
			active = true
		}
	}

	return
}
