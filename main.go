package main

import (
	"gin/controller/auth"
	comment "gin/controller/comment"
	follow "gin/controller/follower"
	post "gin/controller/post"
	profile "gin/controller/profile"
	reply "gin/controller/replycomment"
	"gin/database"
	"gin/middelware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.DataMigration()
	r := gin.Default()
	// r.Use(middelware.AuthRequired())
	r.Static("/static", "./static")

	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
	r.POST("/profilecreate", middelware.AuthRequired(), profile.ProfileCreate)
	r.PUT("/profileupdate/:id", middelware.AuthRequired(), profile.ProfileUpdate)
	r.GET("/profileget", profile.GetProfile)
	r.POST("/postcreate", middelware.AuthRequired(), post.CreatePost)
	r.PATCH("/captionupdate/:id", middelware.AuthRequired(), post.CaptionUpdate)
	r.GET("/postget", middelware.AuthRequired(), post.PostsGet)
	r.DELETE("/postdelete/:id", middelware.AuthRequired(), post.PostsDelete)
	r.POST("/postcomment", middelware.AuthRequired(), middelware.Notification(), comment.PostComment)
	r.PATCH("/updatecomment/:id", middelware.AuthRequired(), comment.UpdateComment)
	r.DELETE("/deletecomment/:id", middelware.AuthRequired(), comment.DeleteComment)
	r.POST("/replycomment", middelware.AuthRequired(), middelware.Notification(), reply.PostReplyComment)
	r.PATCH("/updatereplycomment/:id", middelware.AuthRequired(), reply.UpdateReplyComment)
	r.DELETE("/deletereplycomment/:id", middelware.AuthRequired(), reply.DeleteReplyComment)
	r.POST("/postfollower", middelware.AuthRequired(), follow.PostFollower)
	r.Run()

}
