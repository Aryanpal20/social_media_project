package comment

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

func PostComment(c *gin.Context) {
	var input post.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenid := c.GetFloat64("id")

	comment := post.Comment{Text: input.Text, UserId: int(tokenid), PostId: input.PostId, Createdat: time.Now()}
	database.Database.Create(&comment)
	c.JSON(http.StatusCreated, gin.H{"Comment": comment})

	var posts post.Post
	var comments post.Comment
	var users user.User
	database.Database.Where("user_id = ?", int(tokenid)).Find(&comments)
	// fmt.Println(comments, "vdjjvdvbdhvb")
	database.Database.Where("id = ?", comment.PostId).Find(&posts)
	// fmt.Println(posts, "e;hdvdvdvbj")
	database.Database.Where("id = ?", posts.User_id).Find(&users)
	// fmt.Println(users, "viuuhsfdivgdvvd")
	var from user.User
	database.Database.Where("id = ?", int(tokenid)).Find(&from)
	// fmt.Println(from, "elvhdsvgdjvgdvgdv")
	var notification = NotificationPayload{
		To:           users.Email,
		Notification: from.Email + " commented no your post",
	}
	c.JSON(http.StatusCreated, gin.H{"Data": notification})
}

func UpdateComment(c *gin.Context) {
	var comments post.Comment
	tokenid := c.GetFloat64("id")
	database.Database.Where("user_id = ?", int(tokenid)).Find(&comments)
	if comments.UserId == int(tokenid) {
		if err := database.Database.Where("user_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		text := c.PostForm("text")
		comment := post.Comment{Text: text}
		database.Database.Model(&comments).Updates(comment)
	}

}

func DeleteComment(c *gin.Context) {
	var comments post.Comment
	tokenid := c.GetFloat64("id")
	database.Database.Where("user_id = ?", int(tokenid)).Find(&comments)
	if comments.UserId == int(tokenid) {
		if err := database.Database.Where("user_id = ?", c.Param("id")).First(&comments).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		database.Database.Delete(&comments)
		c.JSON(http.StatusOK, gin.H{"message": "your Post will be deleted successfully"})
	}
}
