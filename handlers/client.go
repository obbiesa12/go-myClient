package handlers

import (
	"go-myClient/config"
	"go-myClient/models"
	"go-myClient/services"
	"go-myClient/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	config *config.Config
	svc    *services.ClientServ
}

func NewClientHandler(cfg *config.Config, s *services.ClientServ) *ClientHandler {
	return &ClientHandler{config: cfg, svc: s}
}

func (h *ClientHandler) Create(c echo.Context) error {
	var req models.Client_Req
	var client models.Client
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	url, err := utils.S3Upload(req.ClientLogoFP, h.config.S3Name, h.config.S3Region)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "file path tidak ditemukan")
	}

	client = models.Client{
		Name:         req.Name,
		Slug:         req.Slug,
		IsProject:    req.IsProject,
		SelfCapture:  req.SelfCapture,
		ClientPrefix: req.ClientPrefix,
		ClientLogo:   url,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		City:         req.City,
	}

	if err := h.svc.Create(&client); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, client)
}

func (h *ClientHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID salah")
	}

	var client models.Client
	var req models.Client_Req
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	url, err := utils.S3Upload(req.ClientLogoFP, h.config.S3Name, h.config.S3Region)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "file path tidak ditemukan")
	}

	client = models.Client{
		ID:           id,
		Name:         req.Name,
		Slug:         req.Slug,
		IsProject:    req.IsProject,
		SelfCapture:  req.SelfCapture,
		ClientPrefix: req.ClientPrefix,
		ClientLogo:   url,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		City:         req.City,
	}
	if err := h.svc.Update(&client); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID tidak sesuai")
	}
	if err := h.svc.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *ClientHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID salah")
	}
	client, err := h.svc.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, client)
}
