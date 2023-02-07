package profilemodel

import "time"

type Profile struct {
	ID         int       `json:"id"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Image      string    `json:"img"`
	UserID     int       `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
