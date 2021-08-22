package organization

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"catalog/internal/app"
	"catalog/internal/domain/building"
	"catalog/internal/domain/category"
	"catalog/internal/domain/organization"
	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/pagination"
)

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
