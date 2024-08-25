// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dbq

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Workerz struct {
	ID        int64            `json:"id"`
	UserID    string           `json:"user_id"`
	FullName  string           `json:"full_name"`
	Email     string           `json:"email"`
	Role      string           `json:"role"`
	Country   string           `json:"country"`
	Salary    float64          `json:"salary"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
