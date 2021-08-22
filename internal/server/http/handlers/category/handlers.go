package category

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
	"catalog/internal/domain/category"
	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/pagination"
)

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
