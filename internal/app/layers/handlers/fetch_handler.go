package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-p3k/internal/app/layers/services"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
	"github.com/iki-rumondor/go-p3k/internal/utils"
)

type FetchHandler struct {
	Service *services.FetchService
}

func NewFetchHandler(service *services.FetchService) *FetchHandler {
	return &FetchHandler{
		Service: service,
	}
}

func (h *FetchHandler) GetGuests(c *gin.Context) {

	resp, err := h.Service.GetGuests()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetGuestByUuid(c *gin.Context) {

	uuid := c.Param("uuid")
	resp, err := h.Service.GetGuestByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetCategories(c *gin.Context) {

	resp, err := h.Service.GetCategories()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetCategoryByUuid(c *gin.Context) {

	uuid := c.Param("uuid")
	resp, err := h.Service.GetCategoryByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetShops(c *gin.Context) {

	resp, err := h.Service.GetShops()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetShopByUuid(c *gin.Context) {

	uuid := c.Param("uuid")
	resp, err := h.Service.GetShopByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetAllProducts(c *gin.Context) {
	resp, err := h.Service.GetAllProducts()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetPublicProductByUuid(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := h.Service.GetPublicProductByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetProducts(c *gin.Context) {
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetProducts(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetProductByUuid(c *gin.Context) {
	userUuid := c.GetString("uuid")
	uuid := c.Param("uuid")
	resp, err := h.Service.GetProductByUuid(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetProductTransactions(c *gin.Context) {
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetProductTransactions(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetProductTransactionByUuid(c *gin.Context) {
	userUuid := c.GetString("uuid")
	uuid := c.Param("uuid")
	resp, err := h.Service.GetProductTransactionByUuid(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}
