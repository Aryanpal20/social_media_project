package comment

import (
	"gin/database"
	comment "gin/models/post_model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	var input comment.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenid := c.GetFloat64("id")

	comment := comment.Comment{Text: input.Text, UserId: int(tokenid), PostId: input.PostId, Createdat: time.Now()}
	database.Database.Create(&comment)
	c.JSON(http.StatusCreated, gin.H{"Comment": comment})
}

func UpdateComment(c *gin.Context) {
	var comments comment.Comment
	tokenid := c.GetFloat64("id")
	database.Database.Where("user_id = ?", int(tokenid)).Find(&comments)
	if comments.UserId == int(tokenid) {
		if err := database.Database.Where("user_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		text := c.PostForm("text")
		comment := comment.Comment{Text: text}
		database.Database.Model(&comments).Updates(comment)
	}

}

func DeleteComment(c *gin.Context) {
	var comments comment.Comment
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
