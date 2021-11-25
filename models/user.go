package models

import "time"

// User structs.
type User struct {
	ID            string    `json:"id" structs:"id" bson:"_id" db:"id"`
	Username      string    `json:"username" structs:"username" bson:"username" db:"username"`
	TokenIssuesAt time.Time `json:"token_issues_at" structs:"token_issues_at" bson:"token_issues_at" db:"token_issues_at"`
}
