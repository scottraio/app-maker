package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	plum "github.com/scottraio/plum"
	memory "github.com/scottraio/plum/memory"
)

type AgentRequest struct {
	Question    string               `json:"question"`
	Version     string               `json:"version"`
	ChatHistory []memory.ChatHistory `json:"chat_history"`
}

func (ar *AgentRequest) Memory() *memory.Memory {
	return memory.LoadMemory(ar.ChatHistory)
}

func Server() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"*"}
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ok": true,
		})
	})

	r.POST("/agent", func(c *gin.Context) {
		var req AgentRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		agent := plum.App.Agents["chef"]
		answer := agent.Remember(req.Memory()).Answer(req.Question)

		response := gin.H{
			"question": req.Question,
			"answer":   answer,
		}

		c.JSON(200, response)
	})

	// Add OPTIONS handler for CORS preflight requests
	r.OPTIONS("/agent", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.Run(":" + plum.App.Port)
}
