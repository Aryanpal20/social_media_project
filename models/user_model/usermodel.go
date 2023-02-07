package usermodel

import (
	follow "gin/models/follower_model"
	post "gin/models/post_model"
	profile "gin/models/profile_model"
	"time"
)

type User struct {
	ID        int               `json:"id"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	CreatedAt time.Time         `json:"created_at"`
	Profiles  []profile.Profile `gorm:"ForeginKey:UserID"`
	Posts     []post.Post       `gorm:"ForeignKey:User_id"`
	Comments  []post.Comment    `gorm:"ForeignKey:UserId"`
	Reply     []post.Reply      `gorm:"ForeignKey:User_ID"`
	Followers []follow.Follower `gorm:"ForeignKey:Following_ID"`
}
