package handler

import (
	"fmt"
	"go_http/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"Message"`
}
type Handler struct {
	storage storage.Storage
}

func NewHandler(storage storage.Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateAuthors(c *gin.Context) {
	var author storage.Author

	if err := c.BindJSON(&author); err != nil {
		fmt.Printf("failed to bind author: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
		return
	}
	h.storage.Create(&author)

	c.JSON(http.StatusOK, Response{Message: "author created"})
}

func (h *Handler) UpdateAuthors(c *gin.Context) {
	var author storage.Author
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Printf("Failed to update author: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
		return
	}

	if err := c.BindJSON(&author); err != nil {
		fmt.Printf("Failed bind author: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
		return
	}
	h.storage.Update(id, &author)

	c.JSON(http.StatusOK, Response{Message: "author updated"})
}

func (h *Handler) GetAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to update author: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
		return
	}
	author, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("Failed to update author: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, author)
}

func (h *Handler) DeleteAuthors(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Printf("Failed to update author: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
		return
	}
	if err := h.storage.Delete(id); err != nil {
		fmt.Printf("Failed to delete author: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, Response{Message: "author created"})
}
