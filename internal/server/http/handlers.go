package internalhttp

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/minipkg/selection_condition"
	"github.com/pkg/errors"

	"catalog/internal/app"
	"catalog/internal/domain/building"
	"catalog/internal/domain/category"
	"catalog/internal/domain/organization"
	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/pagination"
)

func GetBuildingsHandler(c *gin.Context, app *app.App) {
	cond := building.QueryConditions{
		Pagination: pagination.New(pagination.DefaultPage, pagination.DefaultPerPage),
	}
	if err := c.ShouldBindQuery(&cond); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buildings, err := app.Domain.Building.Service.Query(context.Background(), &cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal})
		return
	}

	count, err := app.Domain.Building.Service.Count(context.Background(), &cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": buildings, "count": count})
}

func GetBuildingHandler(c *gin.Context, app *app.App) {
	buildingID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bld, err := app.Domain.Building.Service.Get(context.Background(), uint(buildingID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
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

func GetCategoriesHandler(c *gin.Context, app *app.App) {
	cond := category.QueryConditions{
		Pagination: pagination.New(pagination.DefaultPage, pagination.DefaultPerPage),
	}
	if err := c.ShouldBindQuery(&cond); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categories, err := app.Domain.Category.Service.Query(context.Background(), &cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	count, err := app.Domain.Category.Service.Count(context.Background(), &cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": categories, "count": count})
}

func GetCategoryHandler(c *gin.Context, app *app.App) {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bld, err := app.Domain.Category.Service.Get(context.Background(), uint(categoryID))
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": apperror.ErrNotFound})
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categoryId": buildingID})
}

func DeleteCategoryHandler(c *gin.Context, app *app.App) {
	buildingId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.Domain.Category.Service.Delete(context.Background(), buildingId)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityHasChilds) {
			c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrEntityHasChilds.Error()})
			return
		}

		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetOrganizationsHandler(c *gin.Context, app *app.App) {
	cond := organization.QueryConditions{
		Pagination: pagination.New(pagination.DefaultPage, pagination.DefaultPerPage),
	}
	if err := c.ShouldBindQuery(&cond); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organizations, err := app.Domain.Organization.Service.Query(context.Background(), &cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	count, err := app.Domain.Organization.Service.Count(context.Background(), &cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": organizations, "count": count})
}

func GetOrganizationHandler(c *gin.Context, app *app.App) {
	organizationID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bld, err := app.Domain.Organization.Service.Get(context.Background(), uint(organizationID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"buildingId": buildingID})
}

func DeleteOrganizationHandler(c *gin.Context, app *app.App) {
	organizationID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctgCond := &category.QueryConditions{
		OrganizationID: uint(organizationID),
		Pagination:     pagination.New(pagination.DefaultPage, pagination.DefaultPerPage),
	}
	categories, err := app.Domain.Category.Service.Query(context.Background(), ctgCond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}
	if len(categories) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": apperror.ErrBadRequest.Error(), "categories": categories})
		return
	}

	bldCond := &building.QueryConditions{
		OrganizationID: uint(organizationID),
		Pagination:     pagination.New(pagination.DefaultPage, pagination.DefaultPerPage),
	}
	buildings, err := app.Domain.Building.Service.Query(context.Background(), bldCond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}
	if len(buildings) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": apperror.ErrBadRequest.Error(), "buildings": buildings})
		return
	}

	err = app.Domain.Organization.Service.Delete(context.Background(), uint(organizationID))
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

func UpdateCategory2OrganizationsHandler(c *gin.Context, app *app.App) {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	err = app.Domain.Category.Service.BindOrganizations(context.Background(), uint(categoryID), organizationIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrBadRequest.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
