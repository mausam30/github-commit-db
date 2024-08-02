package api

import (
	"github-commit-viewer/pkg/github"
	"github-commit-viewer/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCommits(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "username is required")
		return
	}

	commits, err := github.FetchCommits(username, "")
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, commits)
}

func GetCommitsByRepo(c *gin.Context) {
	username := c.Query("username")
	repo := c.Param("repo")
	if username == "" || repo == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "username and repo are required")
		return
	}

	commits, err := github.FetchCommits(username, repo)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, commits)
}
