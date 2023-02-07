package database

import (
	"fmt"
	follow "gin/models/follower_model"
	post "gin/models/post_model"
	profile "gin/models/profile_model"
	user "gin/models/user_model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/social_media?parseTime=true"

var err error

func DataMigration() {

	DB, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())

		panic("connection Failed")
	}
	DB.AutoMigrate(user.User{}, profile.Profile{}, post.Post{}, post.Comment{}, post.Reply{}, follow.Follower{})
	Database = DB
}
