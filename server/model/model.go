package model

import "time"

type Workerz struct {
	Id         string    `json:"id"`
	UserId     string    `json:"user_id,omitempty"`
	FullName   string    `json:"full_name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Role       string    `json:"role,omitempty"`
	Country    string    `json:"country,omitempty"`
	Salary     float64   `json:"salary,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateddAt time.Time `json:"updated_at"`
}
