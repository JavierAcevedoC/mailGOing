package api

import (
	"mailGOing/internal/email"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	To      string `json:"to" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/send-email", func(c *gin.Context) {
		var req EmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := email.SendEmail(req.To, req.Subject, req.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully!"})
	})

	r.POST("/send-email-template", func(c *gin.Context) {
		var req EmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := email.SendEmailWithTemplate(req.To, req.Subject, req.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully!"})
	})
	return r
}
