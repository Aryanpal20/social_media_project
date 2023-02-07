package replycomment

import (
	"gin/database"
	comment "gin/models/post_model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostReplyComment(c *gin.Context) {
	var input comment.Reply
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenid := c.GetFloat64("id")

	reply := comment.Reply{Text: input.Text, Comment_ID: input.Comment_ID, User_ID: int(tokenid), Createdat: time.Now()}
	database.Database.Create(&reply)
	c.JSON(http.StatusCreated, gin.H{"Comment": reply})
}

func UpdateReplyComment(c *gin.Context) {
	var replys comment.Reply
	tokenid := c.GetFloat64("id")
	database.Database.Where("userid = ?", int(tokenid)).Find(&replys)
	if replys.User_ID == int(tokenid) {
		if err := database.Database.Where("userid = ?", c.Param("id")).Find(&replys).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		text := c.PostForm("text")
		reply := comment.Comment{Text: text}
		database.Database.Model(&replys).Updates(reply)
	}
}

func DeleteReplyComment(c *gin.Context) {
	var replys comment.Reply
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
