package controllers

import (
	"inverntory-adi-sanbercode/database"
	"inverntory-adi-sanbercode/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetALLTransaksi(c *gin.Context) {
	var transaksiBarang []models.TransaksiBarang
	if err := database.DB.Preload("MasterBarang").Find(&transaksiBarang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": transaksiBarang})

}

func GetTransaksi(c *gin.Context) {
	id := c.Param("id")
	var transaksiBarang []models.TransaksiBarang
	if err := database.DB.
		Preload("MasterBarang").
		Where("barang_id=?", id).
		Find(&transaksiBarang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(transaksiBarang) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Belum ada transaksi untuk barang ini"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"barang_id": id, "transaksi_barang": transaksiBarang})

}

func CreateTransaksiById(c *gin.Context) {
	id := c.Param("id")

	var barang models.MasterBarang
	if err := database.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Barang tidak ditemukan"})
		return
	}

	var input struct {
		Type       string `json:"type" binding:"required"`
		Qty        int    `json:"qty" binding:"required"`
		Keterangan string `json:"keterangan"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaksi := models.TransaksiBarang{
		BarangID:   barang.ID,
		Type:       input.Type,
		Qty:        input.Qty,
		Keterangan: input.Keterangan,
	}

	if err := database.DB.Create(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Preload("MasterBarang").First(&transaksi, transaksi.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat relasi MasterBarang"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaksi barang berhasil ditambahkan",
		"data":    transaksi,
	})
}

func UpdateTransaksi(c *gin.Context) {
	id := c.Param("id")
	var dataTransaksi models.TransaksiBarang
	if err := database.DB.Where("barang_id = ?", id).First(&dataTransaksi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stok untuk barang ini tidak ditemukan"})
		return
	}
	var input struct {
		Type       string `json:"type" binding:"required"`
		Qty        int    `json:"qty" binding:"required"`
		Keterangan string `json:"keterangan"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&dataTransaksi).Updates(models.TransaksiBarang{
		Type:       input.Type,
		Qty:        input.Qty,
		Keterangan: input.Keterangan,
	})

	database.DB.Preload("MasterBarang").First(&dataTransaksi, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaksi barang berhasil diupdate",
		"data":    dataTransaksi,
	})

}
