package followermodel

import "time"

type Follower struct {
	ID           int       `json:"id"`
	UserID       int       `json:"userid"`
	Following_ID int       `json:"followingid"`
	CreatedAt    time.Time `json:"createdat"`
}
