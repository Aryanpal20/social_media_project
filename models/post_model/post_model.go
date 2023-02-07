package postmodel

import "time"

type Post struct {
	ID           int       `json:"id"`
	Upload_Image string    `json:"upload_img"`
	Caption      string    `json:"caption"`
	User_id      int       `json:"userid"`
	Createat     time.Time `json:"created_at"`
	Comments     []Comment `gorm:"ForeignKey:PostId"`
}

type Comment struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Createdat time.Time `json:"created_at"`
	UserId    int       `json:"userid"`
	PostId    int       `json:"postid"`
	Replys    []Reply   `gorm:"ForeignKey:Comment_ID"`
}

type Reply struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	Createdat  time.Time `json:"creation_at"`
	Comment_ID int       `json:"comment_id"`
	User_ID    int       `json:"userid"`
}
