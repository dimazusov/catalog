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

// @Summary get categories
// @Description get categories by params
// @ID get-categories
// @Accept json
// @Produce json
// @Param with_organization query boolean false "with organization"
// @Param per_page query int false "per page"
// @Param page query int false "page"
// @Success 200 {object} CategoryList
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /categories [get]
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

	c.JSON(http.StatusOK, CategoryList{
		Items: categories,
		Count: count,
	})
}

// @Summary get category by id
// @Description get category by id
// @ID get-category-by-id
// @Accept json
// @Produce json
// @Param id path boolean true "id"
// @Success 200 {object} Category
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /category/{id} [get]
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

// @Summary category category
// @Description category category
// @ID category-category
// @Accept json
// @Produce json
// @Param category body category.Category true "updatable category"
// @Success 200 {string} success
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /category [put]
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

// @Summary create category
// @Description create category
// @ID create-category
// @Accept json
// @Produce json
// @Param category body category.Category true "creatable category"
// @Success 200 {object} category.Category
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /category [post]
func CreateCategoryHandler(c *gin.Context, app *app.App) {
	ctg := category.Category{}
	if err := c.ShouldBindJSON(&ctg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryID, err := app.Domain.Category.Service.Create(context.Background(), &ctg)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}
	ctg.ID = categoryID

	c.JSON(http.StatusOK, ctg)
}

// @Summary delete category by id
// @Description delete category by id
// @ID delete-category-by-id
// @Accept json
// @Produce json
// @Param id path boolean true "id"
// @Success 200 {string} success
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /category/{id} [delete]
func DeleteCategoryHandler(c *gin.Context, app *app.App) {
	categoryId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.Domain.Category.Service.Delete(context.Background(), categoryId)
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

// @Summary bind organizations to category
// @Description bind organizations to category
// @ID category-to-organizations
// @Accept json
// @Produce json
// @Param id path boolean true "id"
// @Param categoryIds body []int true "creatable category"
// @Success 200 {string} success
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /category/{id}/organizations [put]
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
