package middelware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin/controller/auth"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(auth.JwtKey), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Token is expired"})
			panic("invalid token")

		} else {
			c.JSON(http.StatusAccepted, gin.H{"message": "Token is valid"})
			claims := token.Claims.(jwt.MapClaims)["id"]
			fmt.Println("token vaild")
			c.Set("id", claims)
			c.Next()

		}
	}
}

type NotificationPayload struct {
	To           string `json:"to"`
	Notification string `json:"notification"`
}
type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func Notification() gin.HandlerFunc {
	return func(c *gin.Context) {

		var notificationPayload NotificationPayload

		// Set up HTTP request to FCM API
		url := "https://accounts.google.com/o/oauth2/auth"
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add required headers for FCM API
		req.Header.Add("Authorization", "AIzaSyDAruyto8bQJeI18giFauD9wz9tqascG4o")
		req.Header.Add("Content-Type", "application/json")

		// Encode notification payload as JSON
		payload, err := json.Marshal(notificationPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Add JSON payload to request body
		req.Body = nopCloser{bytes.NewReader(payload)}
		// Send HTTP request to FCM API
		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
			return
		}

		defer res.Body.Close()

		// Check HTTP response status code
		if res.StatusCode != http.StatusOK {
			c.JSON(res.StatusCode, gin.H{"error": fmt.Sprintf("FCM API responded with status code %d", res.StatusCode)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Notification sent"})
		fmt.Println(notificationPayload.Notification)
		c.Next()
	}
}
