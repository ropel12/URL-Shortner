package handlers

import (
	"net/http"

	"github.com/ropel12/URL-Shortner/entity"
	"github.com/ropel12/URL-Shortner/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "path is required",
		})
	}

	url, err := h.repo.Get(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
	}

	c.Redirect(http.StatusFound, url.LongURL)
}

func (h *URLHandler) Create(c *gin.Context) {
	var url *entity.URL

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid request",
		})
	}

	url, err := h.repo.Create(url.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError, "message": "Success",
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "code": http.StatusOK, "data": url})
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	var url *entity.URL

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid request",
		})
	}

	url, err := h.repo.CreateCustom(url.LongURL, url.ShortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError, "message": "Success",
		})
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Success", "data": url})
}
