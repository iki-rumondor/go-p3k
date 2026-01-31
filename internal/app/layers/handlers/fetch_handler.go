package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-p3k/internal/app/layers/services"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/request"
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

func (h *FetchHandler) GetCitizens(c *gin.Context) {
	resp, err := h.Service.GetCitizens()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetCitizenByUuid(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := h.Service.GetCitizenByUuid(uuid)
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
	limit := c.Query("limit")
	resp, err := h.Service.GetShops(limit)
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
	limit := c.Query("limit")
	shop := c.Query("shop")
	resp, err := h.Service.GetAllProducts(limit, shop)
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

func (h *FetchHandler) GetProductTransactionsByShop(c *gin.Context) {
	is_accept := c.Query("is_accept") == "true"
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetProductTransactionsByShop(userUuid, is_accept)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetMembers(c *gin.Context) {
	group := c.Query("group")
	resp, err := h.Service.GetMembers(group)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetMemberByUuid(c *gin.Context) {

	uuid := c.Param("uuid")
	resp, err := h.Service.GetMemberByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetActivities(c *gin.Context) {
	limit := c.Query("limit")
	group := c.Query("group")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	resp, err := h.Service.GetActivities(limit, group, startDate, endDate)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetActivityByUuid(c *gin.Context) {
	member := c.Query("member")
	queries := request.ActivityQuery{
		Member: member,
	}
	uuid := c.Param("uuid")
	resp, err := h.Service.GetActivityByUuid(uuid, queries)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetMembersNotInActivity(c *gin.Context) {
	activityUuid := c.Param("activityUuid")
	resp, err := h.Service.GetMembersNotInActivity(activityUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetMemberActivity(c *gin.Context) {
	userUuid := c.GetString("uuid")
	activityUuid := c.Param("activityUuid")
	resp, err := h.Service.GetMemberActivity(userUuid, activityUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetAdminDashboard(c *gin.Context) {
	resp, err := h.Service.GetAdminDashboard()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetShopDashboard(c *gin.Context) {
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetShopDashboard(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetGuestDashboard(c *gin.Context) {
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetGuestDashboard(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetShopByUser(c *gin.Context) {
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetShopByUser(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetMemberDashboard(c *gin.Context) {
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetMemberDashboard(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetMemberActivities(c *gin.Context) {
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetMemberActivities(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetMemberByUser(c *gin.Context) {
	userUuid := c.GetString("uuid")
	resp, err := h.Service.GetMemberByUser(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetAllMemberActivities(c *gin.Context) {
	resp, err := h.Service.GetAllMemberActivities()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetProductImage(c *gin.Context) {
	filename := c.Param("filename")
	folder := "internal/files/products"
	pathFile := filepath.Join(folder, filename)
	c.File(pathFile)
}

func (h *FetchHandler) GetActivityImage(c *gin.Context) {
	filename := c.Param("filename")
	folder := "internal/files/activities"
	pathFile := filepath.Join(folder, filename)
	c.File(pathFile)
}

func (h *FetchHandler) GetShopImage(c *gin.Context) {
	filename := c.Param("filename")
	folder := "internal/files/shops"
	pathFile := filepath.Join(folder, filename)
	c.File(pathFile)
}

func (h *FetchHandler) GetIdentityImage(c *gin.Context) {
	filename := c.Param("filename")
	folder := "internal/files/identities"
	pathFile := filepath.Join(folder, filename)
	c.File(pathFile)
}

func (h *FetchHandler) GetTransactionProofImage(c *gin.Context) {
	filename := c.Param("filename")
	folder := "internal/files/transaction_proofs"
	pathFile := filepath.Join(folder, filename)
	c.File(pathFile)
}

func (h *FetchHandler) GetAttendanceImage(c *gin.Context) {
	filename := c.Param("filename")
	folder := "internal/files/attendances"
	pathFile := filepath.Join(folder, filename)
	c.File(pathFile)
}
