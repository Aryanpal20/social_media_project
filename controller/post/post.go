package post

import (
	"fmt"
	"gin/database"
	post "gin/models/post_model"
	entity "gin/models/user_model"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	ID        int       `json:"id"`
	Captions  string    `json:"caption"`
	Img       string    `json:"img_url"`
	Create_AT time.Time `json:"createdat"`
}

func CreatePost(c *gin.Context) {
	var users entity.User
	var message Message
	var posts post.Post
	caption := c.PostForm("caption")
	tokenid := c.GetFloat64("id")
	database.Database.Where("id = ?", int(tokenid)).Find(&users)
	fmt.Println(users)
	if users.ID == int(tokenid) {
		file, header, _ := c.Request.FormFile("img_url")
		filename := header.Filename
		out, err := os.Create("./static/" + filename)
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		scheme1 := "http://"
		if c.Request.TLS != nil {
			scheme := "https://"
			posts.Upload_Image = scheme + c.Request.Host + "./static/" + filename
		}
		posts.Upload_Image = "/static/" + filename
		posts = post.Post{Caption: caption, User_id: int(tokenid), Upload_Image: posts.Upload_Image, Createat: time.Now()}
		Image := scheme1 + c.Request.Host + "/static/" + filename
		message = Message{ID: posts.User_id, Img: Image, Captions: caption}
		database.Database.Create(&posts)
		c.JSON(http.StatusCreated, gin.H{"Profile": message})
	}
}

func CaptionUpdate(c *gin.Context) {
	var posts post.Post
	tokenid := c.GetFloat64("id")
	database.Database.Where("user_id = ?", int(tokenid)).Find(&posts)
	if posts.User_id == int(tokenid) {
		if err := database.Database.Where("user_id = ?", c.Param("id")).First(&posts).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		fmt.Println(posts)
		caption := c.PostForm("caption")
		post := post.Post{Caption: caption}
		fmt.Println(post)
		database.Database.Model(&posts).Updates(post)
		c.JSON(http.StatusAccepted, gin.H{"Data": posts})
	}

}

func PostsGet(c *gin.Context) {
	var posts post.Post
	if err := database.Database.Where("user_id = ?", c.Param("id")).First(&posts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"Data": posts})
}

func PostsDelete(c *gin.Context) {
	var posts post.Post
	tokenid := c.GetFloat64("id")
	database.Database.Where("user_id = ?", int(tokenid)).Find(&posts)
	if posts.User_id == int(tokenid) {
		if err := database.Database.Where("user_id = ?", c.Param("id")).First(&posts).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		database.Database.Delete(&posts)
		c.JSON(http.StatusOK, gin.H{"message": "your Post will be deleted successfully"})
	}
}
