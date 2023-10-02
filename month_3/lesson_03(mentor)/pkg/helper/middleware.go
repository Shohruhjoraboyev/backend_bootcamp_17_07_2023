package helper

import (
	"app/config"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartMiddleware(c *gin.Context) {
	// before request
	fmt.Printf("%s request start path: %s time %v\n", c.Request.Method, c.Request.URL.Path, time.Now())
	c.Next()
}
func EndMiddleware(c *gin.Context) {
	c.Next()
	// after request
	fmt.Printf("%s request end path: %s time %v\n", c.Request.Method, c.Request.URL.Path, time.Now())
}
func LoggerAllInOne(c *gin.Context) {
	// before request
	t := time.Now()
	c.Next()
	// after request
	latency := time.Since(t)

	// access the status we are sending
	status := c.Writer.Status()

	fmt.Printf("Completed %s %s with status code %d in %v\n", c.Request.Method, c.Request.URL.Path, status, latency)
}
func Logger(c *gin.Context) {
	// before request
	beforeRequest(c)
	c.Next()
	// after request
	afterRequest(c)
}
func beforeRequest(c *gin.Context) {
	// before request
	t := time.Now()
	c.Set("start", t)
	c.Next()
}
func afterRequest(c *gin.Context) {
	// Get the start time from the request context
	startTime, exists := c.Get("start")
	if !exists {
		startTime = time.Now()
	}

	// Calculate the request duration
	duration := time.Since(startTime.(time.Time))

	// Log the request completion time and duration
	fmt.Printf("Completed %s %s in %v\n", c.Request.Method, c.Request.URL.Path, duration)
}

// task
func AuthMiddleWare(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    "UNAUTHORIZED!",
			"message": "token not found...",
		})
		return
	}

	staffInfo, err := ParseClaims(token, config.JWTSecretKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    "INVALID TOKEN!",
			"message": "provided token is not valid...",
		})
		return
	}

	c.Set("staff_info", staffInfo)
	c.Next()
}

// password middleware
func ValidatePasswordMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.PostForm("password")
		if !IsValidPassword(password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password. Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, and one digit."})
			c.Abort()
			return
		}
		c.Next()
	}
}
