package internalhttp

import (
	"catalog/internal/domain/category"
	"catalog/internal/domain/organization"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minipkg/pagination"
	"github.com/minipkg/selection_condition"

	"catalog/internal/app"
	"catalog/internal/domain/building"
	"catalog/internal/pkg/apperror"
)

func GetBuildingsHandler(c *gin.Context, app *app.App) {
	cond, err := selection_condition.ParseQueryParams(c.Request.URL.Query(), &building.Building{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := app.Domain.Building.Service.Count(context.Background(), cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal})
		return
	}

	pages := pagination.NewFromRequest(c.Request, int(count))

	cond.Limit = pages.Limit()
	cond.Offset = pages.Offset()

	buildings, err := app.Domain.Building.Service.Query(context.Background(), cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": buildings, "count": count})
}

func GetBuildingHandler(c *gin.Context, app *app.App) {
	buildingId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bld, err := app.Domain.Building.Service.Get(context.Background(), buildingId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetCategoriesHandler(c *gin.Context, app *app.App) {
	cond, err := selection_condition.ParseQueryParams(c.Request.URL.Query(), &building.Building{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := app.Domain.Building.Service.Count(context.Background(), cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal})
		return
	}

	pages := pagination.NewFromRequest(c.Request, int(count))

	cond.Limit = pages.Limit()
	cond.Offset = pages.Offset()

	categories, err := app.Domain.Category.Service.Query(context.Background(), cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": categories, "count": count})
}

func GetCategoryHandler(c *gin.Context, app *app.App) {
	buildingId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bld, err := app.Domain.Category.Service.Get(context.Background(), buildingId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, bld)
}

func UpdateCategoryHandler(c *gin.Context, app *app.App) {
	ctg := category.Category{}
	if err := c.ShouldBindJSON(&ctg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.Domain.Category.Service.Update(context.Background(), &ctg)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func CreateCategoryHandler(c *gin.Context, app *app.App) {
	ctg := category.Category{}
	if err := c.ShouldBindJSON(&ctg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buildingID, err := app.Domain.Category.Service.Create(context.Background(), &ctg)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{"buildingId": buildingID})
}

func DeleteCategoryHandler(c *gin.Context, app *app.App) {
	buildingId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.Domain.Category.Service.Delete(context.Background(), buildingId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetOrganizationsHandler(c *gin.Context, app *app.App) {
	cond, err := selection_condition.ParseQueryParams(c.Request.URL.Query(), &organization.Organization{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := app.Domain.Organization.Service.Count(context.Background(), cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal})
		return
	}

	pages := pagination.NewFromRequest(c.Request, int(count))

	cond.Limit = pages.Limit()
	cond.Offset = pages.Offset()

	organizations, err := app.Domain.Organization.Service.Query(context.Background(), cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": organizations, "count": count})
}

func GetOrganizationHandler(c *gin.Context, app *app.App) {
	buildingId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bld, err := app.Domain.Organization.Service.Get(context.Background(), buildingId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, bld)
}

func UpdateOrganizationHandler(c *gin.Context, app *app.App) {
	org := organization.Organization{}
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.Domain.Organization.Service.Update(context.Background(), &org)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func CreateOrganizationHandler(c *gin.Context, app *app.App) {
	org := organization.Organization{}
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buildingID, err := app.Domain.Organization.Service.Create(context.Background(), &org)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{"buildingId": buildingID})
}

func DeleteOrganizationHandler(c *gin.Context, app *app.App) {
	buildingId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.Domain.Organization.Service.Delete(context.Background(), buildingId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
