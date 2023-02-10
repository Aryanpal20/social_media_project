package replycomment

import (
	"gin/database"
	post "gin/models/post_model"
	user "gin/models/user_model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type NotificationPayload struct {
	To           string `json:"to"`
	Notification string `json:"notification"`
}

func PostReplyComment(c *gin.Context) {
	var input post.Reply
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenid := c.GetFloat64("id")

	reply := post.Reply{Text: input.Text, Comment_ID: input.Comment_ID, User_ID: int(tokenid), Createdat: time.Now()}
	database.Database.Create(&reply)
	c.JSON(http.StatusCreated, gin.H{"Comment": reply})

	var comments post.Comment
	var users user.User
	var replys post.Reply
	database.Database.Where("user_id = ?", int(tokenid)).Find(&replys)
	database.Database.Where("id = ?", replys.Comment_ID).Find(&comments)
	// fmt.Println(comments, "vdjjvdvbdhvb")
	database.Database.Where("id = ?", comments.UserId).Find(&users)
	// fmt.Println(users, "viuuhsfdivgdvvd")
	var from user.User
	database.Database.Where("id = ?", int(tokenid)).Find(&from)
	// fmt.Println(from, "elvhdsvgdjvgdvgdv")
	var notification = NotificationPayload{
		To:           users.Email,
		Notification: from.Email + " replyed to your comment",
	}
	c.JSON(http.StatusCreated, gin.H{"Data": notification})
}

func UpdateReplyComment(c *gin.Context) {
	var replys post.Reply
	tokenid := c.GetFloat64("id")
	database.Database.Where("userid = ?", int(tokenid)).Find(&replys)
	if replys.User_ID == int(tokenid) {
		if err := database.Database.Where("userid = ?", c.Param("id")).Find(&replys).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		text := c.PostForm("text")
		reply := post.Comment{Text: text}
		database.Database.Model(&replys).Updates(reply)
	}
}

func DeleteReplyComment(c *gin.Context) {
	var replys post.Reply
	tokenid := c.GetFloat64("id")
	database.Database.Where("user_id = ?", int(tokenid)).Find(&replys)
	if replys.User_ID == int(tokenid) {
		if err := database.Database.Where("user_id = ?", c.Param("id")).First(&replys).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		database.Database.Delete(&replys)
		c.JSON(http.StatusOK, gin.H{"message": "your Post will be deleted successfully"})
	}
}
