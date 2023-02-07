package middelware

import (
	"fmt"
	"gin/controller/auth"
	"net/http"
	"strconv"
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

const (
	DEFAULT_PAGE_TEXT    = "page"
	DEFAULT_SIZE_TEXT    = "size"
	DEFAULT_PAGE         = "1"
	DEFAULT_PAGE_SIZE    = "10"
	DEFAULT_MIN_PAGESIZE = 10
	DEFAULT_MAX_PAGESIZE = 100
)

func Default() gin.HandlerFunc {
	return New(
		DEFAULT_PAGE_TEXT,
		DEFAULT_SIZE_TEXT,
		DEFAULT_PAGE,
		DEFAULT_PAGE_SIZE,
		DEFAULT_MIN_PAGESIZE,
		DEFAULT_MAX_PAGESIZE,
	)
}

func New(pageText, sizeText, defaultPage, defaultPageSize string, minPageSize, maxPageSize int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the page from the query string and convert it to an integer
		pageStr := c.DefaultQuery(pageText, defaultPage)
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "page number must be an integer"})
			return
		}

		// Validate for positive page number
		if page < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "page number must be positive"})
			return
		}

		// Extract the size from the query string and convert it to an integer
		sizeStr := c.DefaultQuery(sizeText, defaultPageSize)
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "page size must be an integer"})
			return
		}

		// Validate for min and max page size
		if size < minPageSize || size > maxPageSize {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "page size must be between " + strconv.Itoa(minPageSize) + " and " + strconv.Itoa(maxPageSize)})
			return
		}

		// Set the page and size in the gin context
		c.Set(pageText, page)
		c.Set(sizeText, size)

		c.Next()
	}
}
