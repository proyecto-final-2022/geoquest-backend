package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/client"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Client struct {
	service client.Service
}

type questRequest struct {
	Name          string  `json:"name"`
	Qualification float32 `json:"qualification"`
	Description   string  `json:"description"`
	Difficulty    string  `json:"difficulty"`
	Duration      string  `json:"duration"`
	Image         string  `json:"image_url"`
}

type tagRequest struct {
	Description []string `json:"description"`
}

type clientRequest struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func NewClient(s client.Service) *Client {
	return &Client{service: s}
}

// @Summary New client
// @Schemes
// @Description Save new client
// @Tags Clients
// @Accept json
// @Produce json
// @Param quest body clientRequest true "Client to save"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /clients/ [post]
func (u *Client) CreateClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.ClientDTO
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err = u.service.CreateClient(c, req.Name, req.Image); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Clients
// @Schemes
// @Description All clients
// @Tags Clients
// @Accept json
// @Produce json
// @Success 200
// @Failure 422
// @Failure 500
// @Router /clients/ [get]
func (u *Client) GetClients() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		clients, err := u.service.GetClients(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, clients)
	}
}

// @Summary Quests
// @Schemes
// @Description All quests from a client
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /clients/{id}/quests  [get]
func (u *Client) GetClientQuests() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		quests, err := u.service.GetClientQuests(c, paramId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, quests)
	}
}

// @Summary Quests
// @Schemes
// @Description All quests from all clients
// @Tags Clients
// @Accept json
// @Produce json
// @Success 200
// @Failure 422
// @Failure 500
// @Router /clients/quests  [get]
func (u *Client) GetAllQuests() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		quests, err := u.service.GetAllQuests(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, quests)
	}
}

// @Summary New Quest for client
// @Schemes
// @Description Create new quest for a client
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID"
// @Param user body questRequest true "Quest to create for client"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /clients/{id}/quests  [post]
func (u *Client) CreateClientQuest() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req questRequest

		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err = u.service.CreateQuest(c, paramId, req.Name, req.Qualification, req.Description, req.Difficulty, req.Duration, req.Image); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary New Tag for quest
// @Schemes
// @Description Create new tag for a quest
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "Quest ID"
// @Param user body tagRequest true "Tag to add to quest"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /clients/quests/{id}  [post]
func (u *Client) AddTag() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req tagRequest

		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err = u.service.AddTag(c, paramId, req.Description); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}
