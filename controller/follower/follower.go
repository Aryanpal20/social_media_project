package follower

import (
	"fmt"
	"gin/database"
	follow "gin/models/follower_model"
	entity "gin/models/user_model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostFollower(c *gin.Context) {
	var input follow.Follower
	var users entity.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenid := c.GetFloat64("id")
	follower := follow.Follower{UserID: input.UserID}
	database.Database.Where("id = ?", follower.UserID).Find(&users)
	fmt.Println(users)
	if users.ID != int(tokenid) {
		if users.ID == follower.UserID {
			fmt.Println("skcvnjnv")
			follow := follow.Follower{UserID: input.UserID, Following_ID: int(tokenid), CreatedAt: time.Now()}
			database.Database.Create(&follow)
			c.JSON(http.StatusCreated, gin.H{"Comment": follow})
		}
	}

}
