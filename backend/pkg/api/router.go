package api

import (
	"github-commit-viewer/pkg/github"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(github.RateLimiter())

	router.GET("/commits", GetCommits)

	return router
}
