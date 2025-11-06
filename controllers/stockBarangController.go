package controllers

import (
	"inverntory-adi-sanbercode/database"
	"inverntory-adi-sanbercode/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStock(c *gin.Context) {
	var stockBarang []models.StockBarang
	if err := database.DB.Preload("MasterBarang").Find(&stockBarang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stockBarang})

}
func GetStock(c *gin.Context) {
	id := c.Param("id")
	var stockBarang []models.StockBarang
	if err := database.DB.
		Preload("MasterBarang").
		Where("barang_id = ?", id).
		Find(&stockBarang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(stockBarang) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Belum ada stok untuk barang ini"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"barang_id": id, "stok": stockBarang})

}
func CreateStockById(c *gin.Context) {
	id := c.Param("id")

	var barang models.MasterBarang
	if err := database.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Barang tidak ditemukan"})
		return
	}

	var input struct {
		Qty          int  `json:"qty" binding:"required"`
		StatusBarang bool `json:"status_barang"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stock := models.StockBarang{
		BarangID:     barang.ID,
		Qty:          input.Qty,
		StatusBarang: input.StatusBarang,
	}

	if err := database.DB.Create(&stock).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Preload("MasterBarang").First(&stock, stock.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat relasi MasterBarang"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Stok berhasil ditambahkan",
		"data":    stock,
	})
}

func UpdateStock(c *gin.Context) {
	id := c.Param("id")
	var dataStock models.StockBarang
	if err := database.DB.First(&dataStock, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data Stock Tidak ditemukan"})
		return
	}
	var inputStock struct {
		Qty          int  `json:"qty"`
		StatusBarang bool `json:"status_barang"`
	}

	if err := c.ShouldBind(&inputStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&dataStock).Updates(models.StockBarang{
		Qty:          inputStock.Qty,
		StatusBarang: inputStock.StatusBarang,
	})

	database.DB.Preload("MasterBarang").First(&dataStock, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Stock barang berhasil diupdate",
		"data":    dataStock,
	})

}
