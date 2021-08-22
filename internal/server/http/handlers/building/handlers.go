package building

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/minipkg/selection_condition"

	"catalog/internal/app"
	"catalog/internal/domain/building"
	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/pagination"
	"catalog/internal/server/http/api_error"
)

// @Summary get buildings
// @Description get buildings by params
// @ID get-buildings
// @Accept json
// @Produce json
// @Success 200 {object} BuildingsList
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /buildings [get]
func GetBuildingsHandler(c *gin.Context, app *app.App) {
	cond := building.QueryConditions{
		Pagination: pagination.New(pagination.DefaultPage, pagination.DefaultPerPage),
	}
	if err := c.ShouldBindQuery(&cond); err != nil {
		c.JSON(http.StatusBadRequest, api_error.New(err))
		return
	}

	buildings, err := app.Domain.Building.Service.Query(context.Background(), &cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, api_error.New(apperror.ErrInternal))
		return
	}

	count, err := app.Domain.Building.Service.Count(context.Background(), &cond)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusInternalServerError, api_error.New(apperror.ErrInternal))
		return
	}

	c.JSON(http.StatusOK, BuildingsList{
		Items: buildings,
		Count: count,
	})
}

func GetBuildingHandler(c *gin.Context, app *app.App) {
	buildingID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api_error.New(err))
		return
	}

	bld, err := app.Domain.Building.Service.Get(context.Background(), uint(buildingID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, api_error.New(apperror.ErrInternal))
		return
	}

	c.JSON(http.StatusOK, bld)
}

func UpdateBuildingHandler(c *gin.Context, app *app.App) {
	bdg := building.Building{}
	if err := c.ShouldBindJSON(&bdg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.Domain.Building.Service.Update(context.Background(), &bdg)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func CreateBuildingHandler(c *gin.Context, app *app.App) {
	bdg := building.Building{}
	if err := c.ShouldBindJSON(&bdg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buildingID, err := app.Domain.Building.Service.Create(context.Background(), &bdg)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"buildingId": buildingID})
}

func DeleteBuildingHandler(c *gin.Context, app *app.App) {
	buildingId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.Domain.Building.Service.Delete(context.Background(), buildingId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func UpdateBuilding2OrganizationsHandler(c *gin.Context, app *app.App) {
	buildingID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organizationIDs := []uint{}
	err = c.ShouldBindWith(&organizationIDs, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.Domain.Building.Service.BindOrganizations(context.Background(), uint(buildingID), organizationIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrBadRequest.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
