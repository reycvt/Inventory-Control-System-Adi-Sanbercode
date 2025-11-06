package controllers

import (
	"inverntory-adi-sanbercode/database"
	"inverntory-adi-sanbercode/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetALLBarang(c *gin.Context) {
	var barang []models.MasterBarang
	if err := database.DB.Find(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": barang})

}
func GetBarang(c *gin.Context) {
	id := c.Param("id")
	var barang models.MasterBarang
	if err := database.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Barang tidak di temukan"})
		return
	}
}
func CreateBarang(c *gin.Context) {
	var inputBarang models.MasterBarang
	if err := c.ShouldBind(&inputBarang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := database.DB.Create(&inputBarang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"data": inputBarang})

}

func UpdateBarang(c *gin.Context) {
	id := c.Param("id")
	var barang models.MasterBarang
	if err := database.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak di temukan"})
		return
	}
	var inputBarang models.MasterBarang
	if err := c.ShouldBind(&inputBarang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&barang).Updates(inputBarang)
	c.JSON(http.StatusOK, gin.H{"message": "Barang Berhasil di update", "data": barang})
}

func DeleteBarang(c *gin.Context) {
	id := c.Param("id")
	var barang models.MasterBarang
	if err := database.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	var inputBarang models.MasterBarang
	if err := c.ShouldBind(&inputBarang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Delete(&barang)
	c.JSON(http.StatusOK, gin.H{"message": "Barang berhasil di hapus"})
}
