package controllers

import (
    "net/http"
    "flower-shop/internal/services"

    "github.com/gin-gonic/gin"
)

type FlowerController struct {
    FlowerService services.FlowerService
}

func (fc *FlowerController) GetFlowers(c *gin.Context) {
    flowers, err := fc.FlowerService.GetAllFlowers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving flowers"})
        return
    }
    c.JSON(http.StatusOK, flowers)
}
