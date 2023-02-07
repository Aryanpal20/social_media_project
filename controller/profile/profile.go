package profile

import (
	"fmt"
	"gin/database"
	pro "gin/models/profile_model"
	entity "gin/models/user_model"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	ID        int       `json:"id"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Img       string    `json:"img_url"`
	Create_AT time.Time `json:"createdat"`
}
type pagination struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
}

func ProfileCreate(c *gin.Context) {

	var users entity.User
	var message Message
	var profiles []pro.Profile
	fname := c.PostForm("firstname")
	lname := c.PostForm("lastname")
	// userid := c.PostForm("userid")
	// id, err := strconv.Atoi(userid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	tokenid := c.GetFloat64("id")
	profile := pro.Profile{First_Name: fname, Last_Name: lname, UserID: int(tokenid)}
	database.Database.Where("id = ?", profile.UserID).Find(&users)
	database.Database.Where("user_id = ?", profile.UserID).Find(&profiles)
	// fmt.Println(int(tokenid))
	// // t := c.GetInt("id")
	// // fmt.Println(t)
	fmt.Println(len(profiles))
	if len(profiles) == 0 {
		if users.ID == profile.UserID {
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
				profile.Image = scheme + c.Request.Host + "./static/" + filename
			}
			profile.Image = "/static/" + filename
			Image := scheme1 + c.Request.Host + "/static/" + filename
			message = Message{ID: profile.UserID, Fname: fname, Lname: lname, Img: Image, Create_AT: time.Now()}
			profile = pro.Profile{First_Name: fname, Last_Name: lname, UserID: int(tokenid), Image: profile.Image, CreatedAt: time.Now()}
			database.Database.Create(&profile)
			c.JSON(http.StatusCreated, gin.H{"Profile": message})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "User I'd doesn't Match with valid token"})
		}
	}

}

func ProfileUpdate(c *gin.Context) {

	var profile pro.Profile
	var msg Message
	if err := database.Database.Where("user_id = ?", c.Param("id")).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	fname := c.PostForm("fname")
	lname := c.PostForm("lname")

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
		profile.Image = scheme + c.Request.Host + "./static/" + filename
		fmt.Println(profile.Image)
	}
	profile.Image = "/static/" + filename
	profiles := pro.Profile{First_Name: fname, Last_Name: lname, Image: profile.Image, CreatedAt: time.Now()}
	Image := scheme1 + c.Request.Host + "/static/" + filename
	msg = Message{ID: profile.UserID, Fname: fname, Lname: lname, Img: Image, Create_AT: time.Now()}
	database.Database.Model(&profile).Updates(profiles)
	c.JSON(http.StatusOK, gin.H{"data": msg})
}

func GetProfile(c *gin.Context) {
	var profile []pro.Profile
	page := 1
	pagestr := c.Query("page")
	if pagestr != "" {
		page, _ = strconv.Atoi(pagestr)
	}
	offset := (page - 1) * 2

	database.Database.Limit(2).Offset(offset).Find(&profile)
	c.JSON(http.StatusAccepted, gin.H{"Data": profile,
		"pagination": pagination{
			NextPage:     page + 1,
			PreviousPage: page - 1,
			CurrentPage:  page,
		}})

}
