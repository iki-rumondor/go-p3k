package newest

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-p3k/internal/app/database"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
)

func GetAllRegions(c *gin.Context) {
	db := database.GetDB()

	var regions []models.Region
	if err := db.Find(&regions).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": regions})
}

func CreateRegion(c *gin.Context) {
	db := database.GetDB()

	var input models.Region
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&input).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": input})
}

func GetRegionByID(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var region models.Region
	if err := db.First(&region, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "region not found"})
		return
	}

	c.JSON(200, gin.H{"data": region})
}

func UpdateRegion(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var region models.Region
	if err := db.First(&region, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "region not found"})
		return
	}

	var input models.Region
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	region.Name = input.Name

	if err := db.Save(&region).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": region})
}

func DeleteRegion(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	if err := db.Delete(&models.Region{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "region deleted"})
}
