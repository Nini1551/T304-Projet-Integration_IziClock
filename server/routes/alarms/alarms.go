package alarms

import (
	"net/http"
	"server/initializers"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	alarms := route.Group("/alarms")
	{
		alarms.GET("", get_alarms)
		alarms.PUT("/:id", update_alarms)
	}
}

// get_alarms retrieve and return a list of all alarms sorted by ascending ring date
// @Summary Get all alarms
// @Description Retrieve a list of all alarms from DB and return it sorted by ascending ring date
// @Tags Alarms
// @Produce json
// @Success 200 {array} models.Alarm "Alarms send successfully"
// @Failure 500 "Internal Server Error"
// @Router /alarms [get]
func get_alarms(context *gin.Context) {
	var alarms []models.Alarm
	if err := initializers.DB.Order("ring_date asc").Find(&alarms).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, alarms)
}

func update_alarms(context *gin.Context) {
	idParam := context.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alarm ID"})
		return
	}

	var alarm models.Alarm
	if err := initializers.DB.First(&alarm, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Alarm not found"})
		return
	}

	if err := context.ShouldBindJSON(&alarm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&alarm).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, alarm)
}
